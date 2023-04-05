package main

import (
	"eWallet/Banks"
	"eWallet/Banks/MB"
	"eWallet/Init"
	_ "eWallet/Init"
	"encoding/json"
	"github.com/google/uuid"
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
	bank, err := Banks.GetBank("MB")
	if err != nil {
		log.Info("Get bank failed: ", err.Error())
	}
	var requestLinkMB MB.Request_LINK
	requestLinkMB.WalletId = "20220712008"
	requestLinkMB.ActionType = "LINK"
	requestLinkMB.SourceName = "BUI MINH THU"
	requestLinkMB.SourceNumber = "6424328888888"
	requestLinkMB.SourceType = "ACCOUNT"
	requestLinkMB.AuthenType = "SMS"
	requestLinkMB.Mobile = "0987208011"
	requestLinkMB.NationalId = "001199005524"
	requestLinkMB.ResourceId = uuid.New().String()
	reqByte, err := json.Marshal(requestLinkMB)
	if err != nil {
		log.Error("Marshal data error:", requestLinkMB)
	}
	respByte, err := bank.Link(reqByte)
	if err != nil {
		log.Error("MB link error:", err)
	}
	log.Info("Request link response: ", string(respByte))
	stop := make(chan bool)
	<-stop
}
