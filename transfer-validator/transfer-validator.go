package transfer_validator

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

var ShutdownChan chan bool

func InitService(capServiceAddress string) (*grpc.Server, error) {

	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = defaultPort
	}

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Println("failed to listen: %v", err)
		return nil,err;
	}

	service, err := grpc.Dial(capServiceAddress, grpc.WithInsecure())
	capUpdaterClient := capUpdater.NewCapUpdaterServiceClient(service)

	uri := os.Getenv("DB_HOST")
	if uri == "" {
		uri = defaultDBHost
	}

	capReader := services.InitializeReader(uri)
	kafkaClient, err := services.InitializeKafkaClient("kafka-transaction", "localhost:9092");
	if err != nil {
		log.Println("failed to connect to kafka%v", err)
		return nil, err;
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
	println("Validation")
	println(transfer.Origin)
	println(transfer.Destination)
	println(transfer.Amount)
	println(transfer.Type)

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

	transaction := services.TransactionDTO{Date: time.Now()};
	transaction.FromTransfer(transfer);
	downscale := &capUpdater.CapDownscale{
		AccountID: transfer.Origin,
		Value:     transfer.Amount,
	}
	_, err = t.capUpdaterClient.DownscaleCap(context, downscale);
	if err != nil {
		log.Println("failed updating cap %s", err);
		return &transferService.TransferValidation{
			Transfer:  transfer,
			Validated: false,
			Reason:    "Failed updating cap",
		}, err;
	}
	go t.kafkaClient.SendTransaction(&transaction);
	println("Validated")

	return &transferService.TransferValidation{
		Transfer:  transfer,
		Validated: true,
	}, nil
}
