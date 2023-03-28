package SCB

import (
	log "github.com/jeanphorn/log4go"
	"reflect"
)

type SCB struct{}

type Config struct {
	Link       string `json:"Link"`
	User       string `json:"User"`
	Password   string `json:"Password"`
	SecretKey  string `json:"SecretKey"`
	PrivateKey string `json:"PrivateKey"`
	PublicKey  string `json:"PublicKey"`
}

type Empty struct{}

var Conf Config

func (r SCB) Link(data []byte) ([]byte, error) {
	log.Info("Link", reflect.TypeOf(Empty{}).PkgPath())
	return nil, nil
}

func (r SCB) RequestOTP(data []byte) ([]byte, error) {
	log.Info("RequestOTP", reflect.TypeOf(Empty{}).PkgPath())
	return nil, nil
}

func (r SCB) RequestOTPForLink(data []byte) ([]byte, error) {
	log.Info("RequestOTPForLink", reflect.TypeOf(Empty{}).PkgPath())
	return nil, nil
}

func (r SCB) TopupByCard(data []byte) ([]byte, error) {
	log.Info("TopupByCard", reflect.TypeOf(Empty{}).PkgPath())
	return nil, nil
}

func (r SCB) TopupByAccount(data []byte) ([]byte, error) {
	log.Info("TopupByAccount", reflect.TypeOf(Empty{}).PkgPath())
	return nil, nil
}

func (r SCB) PayByCard(data []byte) ([]byte, error) {
	log.Info("PayByCard", reflect.TypeOf(Empty{}).PkgPath())
	return nil, nil
}

func (r SCB) PayByAccount(data []byte) ([]byte, error) {
	log.Info("PayByAccount", reflect.TypeOf(Empty{}).PkgPath())
	return nil, nil
}

func (r SCB) Withdraw(data []byte) ([]byte, error) {
	log.Info("Withdraw", reflect.TypeOf(Empty{}).PkgPath())
	return nil, nil
}

func (r SCB) TransferToInternalCard(data []byte) ([]byte, error) {
	log.Info("TransferToInternalCard", reflect.TypeOf(Empty{}).PkgPath())
	return nil, nil
}

func (r SCB) TransferToInternalAccount(data []byte) ([]byte, error) {
	log.Info("TransferToInternalAccount", reflect.TypeOf(Empty{}).PkgPath())
	return nil, nil
}

func (r SCB) CancelLink(data []byte) ([]byte, error) {
	log.Info("CancelLink", reflect.TypeOf(Empty{}).PkgPath())
	return nil, nil
}

func (r SCB) TransferToExternalCard(data []byte) ([]byte, error) {
	log.Info("TransferToExternalCard", reflect.TypeOf(Empty{}).PkgPath())
	return nil, nil
}

func (r SCB) TransferToExternalAccount(data []byte) ([]byte, error) {
	log.Info("TransferToExternalAccount", reflect.TypeOf(Empty{}).PkgPath())
	return nil, nil
}

func (r SCB) TransferCardless(data []byte) ([]byte, error) {
	log.Info("TransferCardless", reflect.TypeOf(Empty{}).PkgPath())
	return nil, nil
}

func (r SCB) TransferVisaMasterCard(data []byte) ([]byte, error) {
	log.Info("TransferVisaMasterCard", reflect.TypeOf(Empty{}).PkgPath())
	return nil, nil
}

func (r SCB) QueryBalance(data []byte) ([]byte, error) {
	log.Info("QueryBalance", reflect.TypeOf(Empty{}).PkgPath())
	return nil, nil
}

func (r SCB) QueryStatus(data []byte) ([]byte, error) {
	log.Info("QueryStatus", reflect.TypeOf(Empty{}).PkgPath())
	return nil, nil
}

func (r SCB) CheckGuaranteeBlance(data []byte) ([]byte, error) {
	log.Info("CheckGuaranteeBlance", reflect.TypeOf(Empty{}).PkgPath())
	return nil, nil
}

func (r SCB) CheckCustomerInfo(data []byte) ([]byte, error) {
	log.Info("CheckCustomerInfo", reflect.TypeOf(Empty{}).PkgPath())
	return nil, nil
}

func (r SCB) PushNotification(data []byte) ([]byte, error) {
	log.Info("PushNotification", reflect.TypeOf(Empty{}).PkgPath())
	return nil, nil
}
