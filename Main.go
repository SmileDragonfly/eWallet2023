package main

import (
	"eWallet/Banks"
	"log"
)

func main() {
	// Test call
	log.Println("----------Start eWallet----------")
	bank, err := Banks.GetBank("SCB")
	if err != nil {
		log.Println("Get bank failed: ", err.Error())
	}
	bank.Link(nil)
}
