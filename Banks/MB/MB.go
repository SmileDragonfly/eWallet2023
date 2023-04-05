package MB

import (
	"bytes"
	"crypto"
	randCrypto "crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"errors"
	"github.com/google/uuid"
	log "github.com/jeanphorn/log4go"
	"io"
	"math/rand"
	"net/http"
	"os"
)

type MB struct{}

type Config struct {
	Link         string `json:"Link"`
	User         string `json:"User"`
	Password     string `json:"Password"`
	PrivateKey   string `json:"PrivateKey"`
	PublicKey    string `json:"PublicKey"`
	ChannelID    string `json:"ChannelID"`
	Token        string `json:"Token"`
	TimeGetToken int    `json:"TimeGetToken"`
}

type BodyData interface {
}

type DataToSign struct {
	WalletId   string `json:"walletId"`
	ResourceId string `json:"resourceId"`
	Mobile     string `json:"mobile"`
}

var Conf Config

func GetToken() error {
	log.Info("Get Token")
	url := Conf.Link + "/oauth2/v1/token"
	// Marshal body data
	bodyBuf := []byte("grant_type=client_credentials")
	bodyReader := bytes.NewReader(bodyBuf)
	req, err := http.NewRequest("POST", url, bodyReader)
	if err != nil {
		return err
	}
	// Create header info
	req.SetBasicAuth(Conf.User, Conf.Password)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	// Send request to server
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	// Parse response data
	respBodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	var respBodyStruct Response_AuthorizationToken
	err = json.Unmarshal(respBodyBytes, &respBodyStruct)
	if err != nil {
		return err
	}
	Conf.Token = respBodyStruct.Access_token
	log.Info("Get token successfully: ", Conf.Token)
	return nil
}

func (r MB) Link(data []byte) ([]byte, error) {
	url := Conf.Link + "/private/ms-ewallet-partner/v1.1/link/request"
	log.Info("Link: ", url)
	// Parse request data
	var bodyData Request_LINK
	err := json.Unmarshal(data, &bodyData)
	if err != nil {
		return nil, err
	}
	// Create request to send
	req, err := CreateRequest(url, bodyData, "LK", RandStringRunes(8))
	if err != nil {
		return nil, err
	}
	bodyBuf, err := io.ReadAll(req.Body)
	log.Info("Header: ", req.Header)
	log.Info("Body: ", string(bodyBuf))
	// Do request
	resp, err := DoRequest(req)
	if err != nil {
		return nil, err
	}
	// Parse response data
	respBodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	// Unmarshal response
	var responseBody Response_LINK
	err = json.Unmarshal(respBodyBytes, &responseBody)
	if err != nil {
		return nil, err
	}
	log.Info("Link response data: ", responseBody)
	return respBodyBytes, nil
}

func (r MB) LinkConfirm(data []byte) ([]byte, error) {
	log.Info("Link Confirm")

	return nil, nil
}

func (r MB) CancelLink(data []byte) ([]byte, error) {
	log.Info("Cancel Link")
	url := Conf.Link + "/private/ms-ewallet-partner/v1.1/unlink/request"
	var bodyData Request_UNLINK
	bodyData.WalletId = "20220712008"
	bodyData.AuthenType = "SMS"
	bodyData.Mobile = "0987208011"
	req, err := CreateRequest(url, bodyData, "HK", RandStringRunes(8))
	if err != nil {
		return nil, err
	}
	bodyBuf, err := io.ReadAll(req.Body)
	log.Info("Header: ", req.Header)
	log.Info("Body: ", string(bodyBuf))
	resp, err := DoRequest(req)
	if err != nil {
		return nil, err
	}
	// Parse response data
	respBodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Error("Read response body error: ", err.Error())
		return nil, err
	}
	// Unmarshar response
	var responseBody Response_UNLINK
	err = json.Unmarshal(respBodyBytes, &responseBody)
	if err != nil {
		log.Error("Unmarshal response body error: ", err.Error())
		return nil, err
	}
	log.Info("UnLink response data: ", responseBody)
	return nil, nil
}

func (r MB) CancelLinkConfirm(data []byte) ([]byte, error) {
	log.Info("Cancel Link Confirm")
	return nil, nil
}

func (r MB) RequestOTP(data []byte) ([]byte, error) {
	log.Info("RequestOTP")
	return nil, nil
}

func (r MB) RequestOTPForLink(data []byte) ([]byte, error) {
	log.Info("RequestOTPForLink")
	return nil, nil
}

func (r MB) TopupByCard(data []byte) ([]byte, error) {
	log.Info("TopupByCard")
	return nil, nil
}

func (r MB) TopupByAccount(data []byte) ([]byte, error) {
	log.Info("TopupByAccount")
	return nil, nil
}

func (r MB) PayByCard(data []byte) ([]byte, error) {
	log.Info("PayByCard")
	return nil, nil
}

func (r MB) PayByAccount(data []byte) ([]byte, error) {
	log.Info("PayByAccount")
	return nil, nil
}

func (r MB) Withdraw(data []byte) ([]byte, error) {
	log.Info("Withdraw")
	return nil, nil
}

func (r MB) TransferToInternalCard(data []byte) ([]byte, error) {
	log.Info("TransferToInternalCard")
	return nil, nil
}

func (r MB) TransferToInternalAccount(data []byte) ([]byte, error) {
	log.Info("TransferToInternalAccount")
	return nil, nil
}

func (r MB) TransferToExternalCard(data []byte) ([]byte, error) {
	log.Info("TransferToExternalCard")
	return nil, nil
}

func (r MB) TransferToExternalAccount(data []byte) ([]byte, error) {
	log.Info("TransferToExternalAccount")
	return nil, nil
}

func (r MB) TransferCardless(data []byte) ([]byte, error) {
	log.Info("TransferCardless")
	return nil, nil
}

func (r MB) TransferVisaMasterCard(data []byte) ([]byte, error) {
	log.Info("TransferVisaMasterCard")
	return nil, nil
}

func (r MB) QueryBalance(data []byte) ([]byte, error) {
	log.Info("QueryBalance")
	return nil, nil
}

func (r MB) QueryStatus(data []byte) ([]byte, error) {
	log.Info("QueryStatus")
	return nil, nil
}

func (r MB) CheckGuaranteeBlance(data []byte) ([]byte, error) {
	log.Info("CheckGuaranteeBlance")
	return nil, nil
}

func (r MB) CheckCustomerInfo(data []byte) ([]byte, error) {
	log.Info("CheckCustomerInfo")
	return nil, nil
}

func (r MB) PushNotification(data []byte) ([]byte, error) {
	log.Info("PushNotification")
	return nil, nil
}

func CreateRequest(url string, body BodyData, transactionType string, merchantTransactionID string) (*http.Request, error) {
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
	req.Header.Set("Authorization", "Bearer "+Conf.Token)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("clientMessageId", uuid.New().String())
	// Get transactionID: ChannelID + TransactionType + MerchantTransactionID
	// ChannelID: đối tác được cấp khi tích hợp
	transactionID := Conf.ChannelID + transactionType + merchantTransactionID
	req.Header.Set("transactionId", transactionID)
	// Get sinature for header
	dataSign, err := GetDataToSign(bodyBuf)
	if err != nil {
		return nil, err
	}
	sSignature, err := CreateSignatureSHA256RSA(dataSign, Conf.PrivateKey)
	if err != nil {
		return nil, err
	}
	log.Info("Data Sign: ", string(dataSign))
	req.Header.Set("signature", sSignature)
	return req, nil
}

func DoRequest(req *http.Request) (*http.Response, error) {
	// Send request to server
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Error("Do client request error", err.Error())
		return nil, err
	}
	return resp, err
}

// Format dữ liệu ký: walletId + resourceId + mobile
func CreateSignatureSHA256RSA(src []byte, pathPrivateKey string) (string, error) {
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
	// SHA256 src data
	byteSha256 := sha256.Sum256(src)
	// Sign data using private key
	byteSignedData, err := rsa.SignPKCS1v15(randCrypto.Reader, privateKey, crypto.SHA256, byteSha256[:])
	if err != nil {
		return "", nil
	}
	return base64.StdEncoding.EncodeToString(byteSignedData), nil
}

func VerifySignatureSHA256RSA(sSignature string, src []byte, pathPublicKey string) error {
	// Decode base64 signature
	byteSignature, err := base64.StdEncoding.DecodeString(sSignature)
	if err != nil {
		return err
	}
	// SHA256 MD5 data
	byteSha256 := sha256.Sum256(src)
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
	return rsa.VerifyPKCS1v15(rsaPublicKey, crypto.SHA256, byteSha256[:], byteSignature)
}

func GetDataToSign(data []byte) ([]byte, error) {
	// Get fields need to in body
	var structToSign DataToSign
	err := json.Unmarshal(data, &structToSign)
	if err != nil {
		return nil, err
	}
	// Get value only of these fields
	var byteSign []byte
	byteSign = append(byteSign, []byte(structToSign.WalletId)...)
	byteSign = append(byteSign, []byte(structToSign.ResourceId)...)
	byteSign = append(byteSign, []byte(structToSign.Mobile)...)
	return byteSign, nil
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
