package transfer_validator

import (
	"context"
	"flag"
	tfc_cap_updater "github.com/the-final-codedown/tfc-cap-updater/proto/tfc/cap/updater"
	cap_reader "github.com/the-final-codedown/tfc-transfer-validator/cap-reader"
	cap_updater "github.com/the-final-codedown/tfc-transfer-validator/cap-updater"
	transferservice "github.com/the-final-codedown/tfc-transfer-validator/proto"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
)

type TransferValidator struct {
	capUpdaterClient tfc_cap_updater.CapUpdaterServiceClient
	capReader        cap_reader.CapReader
}

const defaultPort string = ":50052"

const defaultDBHost = "mongodb://localhost:27017"

var ShutdownChan chan bool

func InitService(cap_service_adress string) (*grpc.Server, error) {
	flag.Parse()
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = defaultPort
	}

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	capUpdaterClient, err := cap_updater.GetCapStub(cap_service_adress)
	uri := os.Getenv("DB_HOST")
	if uri == "" {
		uri = defaultDBHost
	}
	capReader := cap_reader.InitializeReader(uri)
	validator := &TransferValidator{capUpdaterClient: capUpdaterClient, capReader: *capReader}

	grpcServer := grpc.NewServer()
	transferservice.RegisterTransferValidatorServer(grpcServer, validator)
	ShutdownChan = make(chan bool)
	go func() {
		grpcServer.Serve(lis)
		ShutdownChan <- true
	}()

	return grpcServer, err

}

func (t TransferValidator) Pay(context context.Context, transfer *transferservice.Transfer) (*transferservice.TransferValidation, error) {
	println("Validation")
	cap, err := t.capReader.GetCap(transfer.Origin)
	if err != nil {
		println("Error fetching cap")
		println(err)
		return &transferservice.TransferValidation{
			Transfer:  transfer,
			Validated: false,
			Reason:    "Error fetching cap",
		}, err
	}
	if cap < transfer.Amount {
		return &transferservice.TransferValidation{
			Transfer:  transfer,
			Validated: false,
			Reason:    "Exeeding Cap",
		}, nil
	}
	downscale := &tfc_cap_updater.CapDownscale{
		AccountID: transfer.Origin,
		Delta:     transfer.Amount,
	}
	_, _ = t.capUpdaterClient.DownscaleCap(context, downscale)
	println("Validated")
	return &transferservice.TransferValidation{
		Transfer:  transfer,
		Validated: true,
	}, nil
}
