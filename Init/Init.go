package Init

import (
	"eWallet/Banks/SCB"
	"encoding/json"
	log "github.com/jeanphorn/log4go"
	"os"
	"reflect"
)

type Empty struct{}

func Init() {
	InitSCB()
}

func InitSCB() {
	// Read config file
	data, err := os.ReadFile("./BankConfigs/SCB.json")
	if err != nil {
		log.Info("Init package failed: ", reflect.TypeOf(Empty{}).PkgPath(), err.Error())
		return
	}
	// Parse json data
	err = json.Unmarshal(data, &SCB.Conf)
	if err != nil {
		log.Info("Init package failed: ", reflect.TypeOf(Empty{}).PkgPath(), err.Error())
		return
	}
	log.Info("Init package successfully: ", reflect.TypeOf(Empty{}).PkgPath(), SCB.Conf)
}
