package main

import (
	"context"
	transferservice "github.com/the-final-codedown/tfc-transfer-validator/proto"
	"google.golang.org/grpc"
)

func main() {
	basic_paying_client()
}
func basic_paying_client() {
	// use the generated client stub
	service, err := grpc.Dial("localhost:50052", grpc.WithInsecure())
	if err != nil {
		println(err)
	}
	cl := transferservice.NewTransferValidatorClient(service)
	_, err = cl.Pay(context.TODO(), &transferservice.Transfer{
		Origin:      1,
		Destination: 0,
		Amount:      400,
	})
	if err != nil {
		println("error", err)
	}
}
