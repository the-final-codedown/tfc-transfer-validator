package main

import (
	"context"
	transferService "github.com/the-final-codedown/tfc-transfer-validator/proto"
	"google.golang.org/grpc"
)

func main() {
	basicPayingClient()
}

func basicPayingClient() {
	// use the generated client stub
	service, err := grpc.Dial("localhost:50052", grpc.WithInsecure())
	if err != nil {
		println(err)
	}
	cl := transferService.NewTransferValidatorServiceClient(service)
	_, err = cl.Pay(context.TODO(), &transferService.Transfer{
		Origin:      1,
		Destination: 0,
		Amount:      400,
	})
	if err != nil {
		println("error", err)
	}
}
