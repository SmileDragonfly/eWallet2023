package Init

import (
	"eWallet/Banks/MB"
	"eWallet/Banks/SCB"
	"encoding/json"
	log "github.com/jeanphorn/log4go"
	"math/rand"
	"os"
	"time"
)

type Empty struct{}

func Init() {
	rand.Seed(time.Now().UnixNano())
	InitSCB()
	InitMB()
}

func InitSCB() {
	// Read config file
	data, err := os.ReadFile("./BankConfigs/SCB/SCB.json")
	if err != nil {
		log.Info("Init package failed: ", err.Error())
		return
	}
	// Parse json data
	err = json.Unmarshal(data, &SCB.Conf)
	if err != nil {
		log.Info("Init package failed: ", err.Error())
		return
	}
	log.Info("Init package successfully: ", SCB.Conf)
}

func InitMB() {
	// Read config file
	data, err := os.ReadFile("./BankConfigs/MB/MB.json")
	if err != nil {
		log.Info("Init package failed: ", err.Error())
		return
	}
	// Parse json data
	err = json.Unmarshal(data, &MB.Conf)
	if err != nil {
		log.Info("Init package failed: ", err.Error())
		return
	}
	// Get token
	if err := MB.GetToken(); err != nil {
		log.Info("Get MB token error", err.Error())
	}
	ticker := time.NewTicker(time.Duration(MB.Conf.TimeGetToken) * time.Second)
	done := make(chan bool)
	go func() {
		for {
			select {
			case <-done:
				return
			case <-ticker.C:
				MB.GetToken()
			}
		}
	}()
	log.Info("Init package successfully: ", MB.Conf)
}
