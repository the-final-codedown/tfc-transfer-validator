package transfervalidator

import (
	"context"
	capUpdater "github.com/the-final-codedown/tfc-cap-updater/proto/tfc/cap/updater"
	transferService "github.com/the-final-codedown/tfc-transfer-validator/proto"
	services "github.com/the-final-codedown/tfc-transfer-validator/service"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"time"
)

type TransferValidator struct {
	capUpdaterClient capUpdater.CapUpdaterServiceClient
	capReader        services.CapReader
	kafkaClient      services.KafkaClient
}

const defaultPort string = ":50052"

const defaultDBHost = "mongodb://localhost:27017"
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

	uri := os.Getenv("DB_HOST")
	if uri == "" {
		uri = defaultDBHost
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
		server.Serve(lis)
		ShutdownChan <- true
	}()

	return server, err

}

func (t TransferValidator) Pay(context context.Context, transfer *transferService.Transfer) (*transferService.TransferValidation, error) {
	println("Payment validation")
	paymentCap, err := t.capReader.GetCap(transfer.Origin)

	if err != nil {
		println("Error fetching cap")
		println(err)
		return &transferService.TransferValidation{
			Transfer:  transfer,
			Validated: false,
			Reason:    "Error fetching cap",
		}, err
	}

	if paymentCap < transfer.Amount {
		return &transferService.TransferValidation{
			Transfer:  transfer,
			Validated: false,
			Reason:    "Exceeding Cap",
		}, nil
	}

	transaction := services.TransactionDTO{Date: time.Now()}
	transaction.FromTransfer(transfer)
	err = t.kafkaClient.SendTransaction(&transaction)
	if err != nil {
		log.Println("error sending to main app", err)
	}
	downscale := &capUpdater.CapDownscale{
		AccountID: transfer.Origin,
		Value:     transfer.Amount,
	}
	resp, err := t.capUpdaterClient.DownscaleCap(context, downscale)
	if err != nil {
		log.Println(err)
	} else {
		log.Println(resp.Accepted)
	}

	println("Payment validated")

	return &transferService.TransferValidation{
		Transfer:  transfer,
		Validated: true,
	}, nil
}
