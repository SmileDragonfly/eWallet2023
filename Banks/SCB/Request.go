package SCB

import (
	"bytes"
	"crypto"
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/x509"
	"eWallet/Utils"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"errors"
	"github.com/google/uuid"
	log "github.com/jeanphorn/log4go"
	"net/http"
	"os"
	"time"
)

type RequestData interface {
}

type RequestBody struct {
	Data            string `json:"Data"`            // Encrypt Triple des using secret key
	FunctionName    string `json:"FunctionName"`    // Each api have it own functional name
	RequestDateTime string `json:"RequestDateTime"` // yyyy-MM-ddTHH:mm:ssZ
	RequestID       string `json:"RequestID"`       // guid
}

type DataOTP struct {
	CustomerID     string `json:"CustomerID"`
	SubscriptionID string `json:"SubscriptionID"`
	RefNumber      string `json:"RefNumber"`
}

func CreateRequestBody(funcName string, reqData RequestData, reqBody *RequestBody) error {
	// Encrypt request data using secret key
	reqDataByte, err := json.Marshal(reqData)
	if err != nil {
		return err
	}
	// Encrypt data
	sEncryptData, err := Utils.EncryptTripleDESB64(string(reqDataByte), Conf.SecretKey)
	if err != nil {
		return err
	}
	reqBody.Data = sEncryptData
	// Assign function name
	reqBody.FunctionName = funcName
	// Get current time: yyyy-MM-ddTHH:mm:ssZ
	now := time.Now()
	strNow := now.UTC().Format("2006-01-02T15:04:05Z")
	reqBody.RequestDateTime = strNow
	// Generate request id: GUID
	guid := uuid.New()
	reqBody.RequestID = guid.String()
	return nil
}

func DoRequest(url string, body BodyData) (*http.Response, error) {
	// Marshal body data
	bodyBuf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader := bytes.NewReader(bodyBuf)
	req, err := http.NewRequest("POST", url, bodyReader)
	if err != nil {
		return nil, err
	}
	// Create header info
	req.SetBasicAuth(Conf.User, Conf.Password)
	sSignature, err := CreateSignatureDataB64(bodyBuf, Conf.PrivateKey)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Signature", sSignature)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	// Send request to server
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Error("Do client request error")
		return nil, err
	}
	return resp, err
}

func CreateSignatureDataB64(src []byte, pathPrivateKey string) (string, error) {
	bytePrivateKey, err := os.ReadFile(pathPrivateKey)
	if err != nil {
		return "", err
	}
	// Load private key
	pemBlock, _ := pem.Decode(bytePrivateKey)
	if pemBlock == nil || pemBlock.Type != "RSA PRIVATE KEY" {
		return "", errors.New("Decode pem file failed")
	}
	privateKey, err := x509.ParsePKCS1PrivateKey(pemBlock.Bytes)
	if err != nil {
		return "", err
	}
	// MD5 body data
	byteMD5 := md5.Sum(src)
	// SHA1 MD5 data
	byteSha1 := sha1.Sum(byteMD5[:])
	// Sign data using private key
	byteSignedData, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA1, byteSha1[:])
	if err != nil {
		return "", nil
	}
	return base64.StdEncoding.EncodeToString(byteSignedData), nil
}

func VerifySignatureB64(sSignature string, src []byte, pathPublicKey string) error {
	// Decode base64 signature
	byteSignature, err := base64.StdEncoding.DecodeString(sSignature)
	if err != nil {
		return err
	}
	// MD5 body data
	byteMD5 := md5.Sum(src)
	// SHA1 MD5 data
	byteSha1 := sha1.Sum(byteMD5[:])
	// Load public key
	pubKeyByte, err := os.ReadFile(pathPublicKey)
	if err != nil {
		return err
	}
	pemBlock, _ := pem.Decode(pubKeyByte)
	if pemBlock == nil || pemBlock.Type != "PUBLIC KEY" {
		return err
	}
	publicKey, err := x509.ParsePKIXPublicKey(pemBlock.Bytes)
	if err != nil {
		return err
	}
	rsaPublicKey, _ := publicKey.(*rsa.PublicKey)
	return rsa.VerifyPKCS1v15(rsaPublicKey, crypto.SHA1, byteSha1[:], byteSignature)
}
