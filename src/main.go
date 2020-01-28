package main

import (
	_ "github.com/jnewmano/grpc-json-proxy/codec"
	"log"
	"os"
)

const defaultCapServiceAddress string = "localhost:50051"

func main() {
	//basic_paying_client()
	capServiceAddress := os.Getenv("CAP_SERVICE_ADDRESS")
	if capServiceAddress == "" {
		capServiceAddress = defaultCapServiceAddress
	}
	_, err := InitService(capServiceAddress)
	if err != nil {
		log.Println("Failed starting services", err)
	} else {
		log.Println("Transfer validator started")
		stopped := <-ShutdownChan
		if stopped {
			log.Println("Gracefully stopped")
		}
	}
}
