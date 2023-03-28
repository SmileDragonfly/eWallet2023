package SCB

import (
	"encoding/json"
	"log"
	"os"
	"reflect"
)

type SCB struct{}

type config struct {
	Link       string `json:"Link"`
	User       string `json:"User"`
	Password   string `json:"Password"`
	SecretKey  string `json:"SecretKey"`
	PrivateKey string `json:"PrivateKey"`
	PublicKey  string `json:"PublicKey"`
}

type Empty struct{}

var cfg config

func init() {
	// Read config file
	data, err := os.ReadFile("./BankConfigs/SCB.json")
	if err != nil {
		log.Println("Init package failed: ", reflect.TypeOf(Empty{}).PkgPath(), err.Error())
		return
	}
	// Parse json data
	err = json.Unmarshal(data, &cfg)
	if err != nil {
		log.Println("Init package failed: ", reflect.TypeOf(Empty{}).PkgPath(), err.Error())
		return
	}
	log.Println("Init package successfully: ", reflect.TypeOf(Empty{}).PkgPath(), cfg)
}

func (r SCB) Link(data []byte) ([]byte, error) {
	log.Println("Link", reflect.TypeOf(Empty{}).PkgPath())
	return nil, nil
}

func (r SCB) RequestOTP(data []byte) ([]byte, error) {
	log.Println("RequestOTP", reflect.TypeOf(Empty{}).PkgPath())
	return nil, nil
}

func (r SCB) RequestOTPForLink(data []byte) ([]byte, error) {
	log.Println("RequestOTPForLink", reflect.TypeOf(Empty{}).PkgPath())
	return nil, nil
}

func (r SCB) TopupByCard(data []byte) ([]byte, error) {
	log.Println("TopupByCard", reflect.TypeOf(Empty{}).PkgPath())
	return nil, nil
}

func (r SCB) TopupByAccount(data []byte) ([]byte, error) {
	log.Println("TopupByAccount", reflect.TypeOf(Empty{}).PkgPath())
	return nil, nil
}

func (r SCB) PayByCard(data []byte) ([]byte, error) {
	log.Println("PayByCard", reflect.TypeOf(Empty{}).PkgPath())
	return nil, nil
}

func (r SCB) PayByAccount(data []byte) ([]byte, error) {
	log.Println("PayByAccount", reflect.TypeOf(Empty{}).PkgPath())
	return nil, nil
}

func (r SCB) Withdraw(data []byte) ([]byte, error) {
	log.Println("Withdraw", reflect.TypeOf(Empty{}).PkgPath())
	return nil, nil
}

func (r SCB) TransferToInternalCard(data []byte) ([]byte, error) {
	log.Println("TransferToInternalCard", reflect.TypeOf(Empty{}).PkgPath())
	return nil, nil
}

func (r SCB) TransferToInternalAccount(data []byte) ([]byte, error) {
	log.Println("TransferToInternalAccount", reflect.TypeOf(Empty{}).PkgPath())
	return nil, nil
}

func (r SCB) CancelLink(data []byte) ([]byte, error) {
	log.Println("CancelLink", reflect.TypeOf(Empty{}).PkgPath())
	return nil, nil
}

func (r SCB) TransferToExternalCard(data []byte) ([]byte, error) {
	log.Println("TransferToExternalCard", reflect.TypeOf(Empty{}).PkgPath())
	return nil, nil
}

func (r SCB) TransferToExternalAccount(data []byte) ([]byte, error) {
	log.Println("TransferToExternalAccount", reflect.TypeOf(Empty{}).PkgPath())
	return nil, nil
}

func (r SCB) TransferCardless(data []byte) ([]byte, error) {
	log.Println("TransferCardless", reflect.TypeOf(Empty{}).PkgPath())
	return nil, nil
}

func (r SCB) TransferVisaMasterCard(data []byte) ([]byte, error) {
	log.Println("TransferVisaMasterCard", reflect.TypeOf(Empty{}).PkgPath())
	return nil, nil
}

func (r SCB) QueryBalance(data []byte) ([]byte, error) {
	log.Println("QueryBalance", reflect.TypeOf(Empty{}).PkgPath())
	return nil, nil
}

func (r SCB) QueryStatus(data []byte) ([]byte, error) {
	log.Println("QueryStatus", reflect.TypeOf(Empty{}).PkgPath())
	return nil, nil
}

func (r SCB) CheckGuaranteeBlance(data []byte) ([]byte, error) {
	log.Println("CheckGuaranteeBlance", reflect.TypeOf(Empty{}).PkgPath())
	return nil, nil
}

func (r SCB) CheckCustomerInfo(data []byte) ([]byte, error) {
	log.Println("CheckCustomerInfo", reflect.TypeOf(Empty{}).PkgPath())
	return nil, nil
}

func (r SCB) PushNotification(data []byte) ([]byte, error) {
	log.Println("PushNotification", reflect.TypeOf(Empty{}).PkgPath())
	return nil, nil
}
