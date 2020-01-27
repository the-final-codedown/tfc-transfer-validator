package main

import (
	_ "github.com/jnewmano/grpc-json-proxy/codec"
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
	_, err := transfervalidator.InitService(capServiceAddress)
	if err != nil {
		println("Failed starting service", err)
	} else {
		println("Transfer validator started")
		stopped := <-transfervalidator.ShutdownChan
		if stopped {
			println("Gracefully stopped")
		}
	}
}
