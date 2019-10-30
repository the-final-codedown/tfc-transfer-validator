package main

import (
	"context"
	transferservice "github.com/the-final-codedown/tfc-transfer-validator/proto"
	transfer_validator "github.com/the-final-codedown/tfc-transfer-validator/transfer-validator"
	"google.golang.org/grpc"
	"os"
)

const default_cap_service_adress string = "localhost:50051"

func main() {
	//basic_paying_client()
	cap_service_adress := os.Getenv("CAP_SERVICE_ADRESS")
	if cap_service_adress == "" {
		cap_service_adress = default_cap_service_adress
	}
	_, err := transfer_validator.InitService(cap_service_adress)
	if err != nil {
		println("Failed starting service", err)
	} else {
		println("Transfer validator started")
		stopped := <-transfer_validator.ShutdownChan
		if stopped {
			println("Gracefuly stopped")
		}
	}
}
