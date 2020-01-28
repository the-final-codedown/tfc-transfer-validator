package main

import (
	capUpdater "github.com/the-final-codedown/tfc-cap-updater/proto"
	transferService "github.com/the-final-codedown/tfc-transfer-validator/proto"
	"github.com/the-final-codedown/tfc-transfer-validator/services"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
)

type TransferValidator struct {
	capUpdaterClient capUpdater.CapUpdaterServiceClient
	capReader        services.CapReader
	kafkaClient      services.KafkaClient
}

const defaultPort string = ":50052"
const defaultHost = "mongodb://localhost:27017"
const defaultKafkaHost = "localhost:9092"

var ShutdownChan chan bool

func InitService(capServiceAddress string) (*grpc.Server, error) {

	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = defaultPort
	}

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Println("failed to listen: %v", err)
		return nil, err
	}

	service, err := grpc.Dial(capServiceAddress, grpc.WithInsecure())
	capUpdaterClient := capUpdater.NewCapUpdaterServiceClient(service)

	uri := ""
	if os.Getenv("DB_HOST") != "" && os.Getenv("DB_PORT") != "" {
		uri = "mongodb://" + os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT")
	} else {
		uri = defaultHost
	}
	capReader := services.InitializeReader(uri)

	uri = os.Getenv("KAFKA_HOST")
	if uri == "" {
		uri = defaultKafkaHost
	}
	kafkaClient, err := services.InitializeKafkaClient("kafka-transaction", uri)
	if err != nil {
		log.Println("failed to connect to kafka%v", err)
		return nil, err
	}

	validator := &TransferValidator{capUpdaterClient: capUpdaterClient, capReader: *capReader, kafkaClient: *kafkaClient}

	server := grpc.NewServer()
	transferService.RegisterTransferValidatorServiceServer(server, validator)

	ShutdownChan = make(chan bool)
	go func() {
		_ = server.Serve(lis)
		ShutdownChan <- true
	}()

	return server, err
}
