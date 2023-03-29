package main

import (
	"eWallet/Banks"
	"eWallet/Init"
	_ "eWallet/Init"
	log "github.com/jeanphorn/log4go"
)

func main() {
	// Init log file
	// load config file, it's optional
	// or log.LoadConfiguration("./example.json", "json")
	// config file could be json or xml
	log.LoadConfiguration("./log4go.json")
	log.Info("----------Init logger----------")
	defer log.Close()
	Init.Init()
	// Test call
	log.Info("----------Start eWallet----------")
	bank, err := Banks.GetBank("SCB")
	if err != nil {
		log.Info("Get bank failed: ", err.Error())
	}
	bank.Link(nil)
	bank.TopupByCard(nil)
	bank.TransferToInternalAccount(nil)
}
