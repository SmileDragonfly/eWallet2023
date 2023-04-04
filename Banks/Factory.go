package Banks

import (
	"eWallet/Banks/MB"
	"eWallet/Banks/SCB"
	"eWallet/Interfaces"
	"errors"
)

func GetBank(name string) (Interfaces.TransactionInterface, error) {
	if name == "SCB" {
		ret := SCB.SCB{}
		return ret, nil
	}
	if name == "MB" {
		ret := MB.MB{}
		return ret, nil
	}
	return nil, errors.New("Invalid bank")
}
