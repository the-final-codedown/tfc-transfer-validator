package test

import (
	"context"
	"github.com/the-final-codedown/tfc-transfer-validator"
	transferService "github.com/the-final-codedown/tfc-transfer-validator/proto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"log"
	"os"
	"testing"
	"time"
)

const (
	mongoAddress string = "mongodb://localhost:27017"
)

func setup() {

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoAddress))
	if err != nil {
		log.Panic(err)
	}
	capCollection := client.Database("tfc").Collection("cap")
	accountOneCap := bson.D{
		{"accountid", 1},
		{"value", 300},
	}
	_, _ = capCollection.InsertOne(context.Background(), accountOneCap)
	client.Disconnect(context.Background())
}

func teardown() {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoAddress))
	if err != nil {
		log.Panic(err)
	}
	capCollection := client.Database("tfc").Collection("cap")
	filter := bson.D{{"accountid", 1}}
	_, _ = capCollection.DeleteMany(context.Background(), filter)
	client.Disconnect(context.Background())
}

/**
* Necessary to run a before all and after all
 */
func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	teardown()
	os.Exit(code)
}

func TestServerStart(t *testing.T) {
	server, err := main.InitService("localhost:50051")
	if err != nil {
		t.Log("Error Creating Server")
		t.Fail()
	}
	defer server.GracefulStop()
	service, err := grpc.Dial("localhost:50052", grpc.WithInsecure())
	if err != nil {
		t.Log("Failed server connection")
		t.Fail()
	}

	// use the generated client stub
	cl := transferService.NewTransferValidatorServiceClient(service)
	if cl == nil {
		t.Log("Failed client creation")
		t.Fail()
	}
	_, err = cl.Pay(context.TODO(), &transferService.Transfer{
		Origin:      "1",
		Destination: "0",
		Amount:      0,
	})
	if err != nil {
		t.Log("Could not reach started server", err)
		t.Fail()
	}
}

func TestServerCanPay(t *testing.T) {
	server, _ := main.InitService("localhost:50051")
	client, _ := initialiseClient()
	defer server.GracefulStop()

	transfer := transferService.Transfer{
		Origin:      "1",
		Destination: "2",
		Amount:      100,
		Type:        transferService.Transfer_CARD,
	}

	result, err := client.Pay(context.Background(), &transfer)

	if err != nil {
		t.Log("failed the paying query", err)
		t.Fail()
	}

	assertTrue(t, result.Validated, "The transfer should be validated", func() { t.Log(result.Reason) })
}

func TestClientCantPay(t *testing.T) {
	server, _ := main.InitService("localhost:50051")
	client, _ := initialiseClient()
	defer server.GracefulStop()

	transfer := transferService.Transfer{
		Origin:      "1",
		Destination: "2",
		Amount:      10000,
		Type:        transferService.Transfer_CARD,
	}
	result, err := client.Pay(context.Background(), &transfer)
	if err != nil {
		t.Log("failed the paying query", err)
		t.Fail()
	}

	assertTrue(t, !result.Validated, "The Transfer Should not be validated", nil)
}

/*
	redefined a utils true assertion function
*/
func assertTrue(t *testing.T, actual bool, message string, onFail func()) {
	if !actual {
		t.Log(message)
		if onFail != nil {
			onFail()
		}
		t.Fail()
	}
}

func initialiseClient() (transferService.TransferValidatorServiceClient, error) {
	service, err := grpc.Dial("localhost:50052", grpc.WithInsecure())

	// use the generated client stub
	cl := transferService.NewTransferValidatorServiceClient(service)
	return cl, err
}
