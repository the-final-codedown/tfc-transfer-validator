package test

import (
	"context"
	transferservice "github.com/the-final-codedown/tfc-transfer-validator/proto"
	transfer_validator "github.com/the-final-codedown/tfc-transfer-validator/transfer-validator"
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
	test_mongo_adress string = "mongodb://localhost:27017"
)

func setup() {

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(test_mongo_adress))
	if err != nil {
		log.Panic(err)
	}
	capCollection := client.Database("tfc").Collection("cap")
	account_one_cap := bson.D{
		{"accountID", 1},
		{"value", 300},
	}
	_, _ = capCollection.InsertOne(context.Background(), account_one_cap)
	client.Disconnect(context.Background())
}

func teardown() {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(test_mongo_adress))
	if err != nil {
		log.Panic(err)
	}
	capCollection := client.Database("tfc").Collection("cap")
	filter := bson.D{{"accountID", 1}}
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
	server, err := transfer_validator.InitService("localhost:50051")
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
	cl := transferservice.NewTransferValidatorClient(service)
	if cl == nil {
		t.Log("Failed client creation")
		t.Fail()
	}
	_, err = cl.Pay(context.TODO(), &transferservice.Transfer{
		Origin:      1,
		Destination: 0,
		Amount:      0,
	})
	if err != nil {
		t.Log("Could not reach started server", err)
		t.Fail()
	}
}

func TestServerCanPay(t *testing.T) {
	server, _ := transfer_validator.InitService("localhost:50051")
	client, _ := initialiseClient()
	defer server.GracefulStop()

	transfer := transferservice.Transfer{
		Origin:      1,
		Destination: 2,
		Amount:      100,
		Type:        transferservice.Transfer_CARD,
	}

	result, err := client.Pay(context.Background(), &transfer)

	if err != nil {
		t.Log("failed the paying query", err)
		t.Fail()
	}

	assertTrue(t, result.Validated, "The transfer should be validated", func() { t.Log(result.Reason) })
}

func TestClientCantPay(t *testing.T) {
	server, _ := transfer_validator.InitService("localhost:50051")
	client, _ := initialiseClient()
	defer server.GracefulStop()

	transfer := transferservice.Transfer{
		Origin:      1,
		Destination: 2,
		Amount:      10000,
		Type:        transferservice.Transfer_CARD,
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
func assertTrue(t *testing.T, actual bool, message string, on_fail func()) {
	if !actual {
		t.Log(message)
		if on_fail != nil {
			on_fail()
		}
		t.Fail()
	}
}
func initialiseClient() (transferservice.TransferValidatorClient, error) {
	service, err := grpc.Dial("localhost:50052", grpc.WithInsecure())

	// use the generated client stub
	cl := transferservice.NewTransferValidatorClient(service)
	return cl, err
}
