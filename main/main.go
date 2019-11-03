package main

import (
	"github.com/the-final-codedown/tfc-transfer-validator/transfer-validator"
	"os"
)

const defaultCapServiceAddress string = "localhost:50051"

func main() {
	//basic_paying_client()
	capServiceAddress := os.Getenv("CAP_SERVICE_ADDRESS")
	if capServiceAddress == "" {
		capServiceAddress = defaultCapServiceAddress
	}
	_, err := transfer_validator.InitService(capServiceAddress)
	if err != nil {
		println("Failed starting service", err)
	} else {
		println("Transfer validator started")
		stopped := <-transfer_validator.ShutdownChan
		if stopped {
			println("Gracefully stopped")
		}
	}
}
