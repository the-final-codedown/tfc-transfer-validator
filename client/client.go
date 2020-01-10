package main

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	transferService "github.com/the-final-codedown/tfc-transfer-validator/proto"
	services "github.com/the-final-codedown/tfc-transfer-validator/service"
	"google.golang.org/grpc"
	"log"
	"time"
)

/**
boiler plate testing package,
Should update the origin and target ids on trial for real accounts ids

*/
func main() {
	originid := "5dc31fad8bc1bb0001d4d853"
	destinationId := "5dc31fad8bc1bb0001d4d854"
	basicPayingClient(originid, destinationId)
}

func basicKafkaPublisher() {

	kafkaClient, err := services.InitializeKafkaClient("kafka-transaction", "localhost:9092")
	if err != nil {
		log.Println("Failed creating kafka client %s", err)
	}
	transaction := services.TransactionDTO{
		Source:   "5dc31dd2adf81c0001fbdf2a",
		Receiver: "5dc31dd2adf81c0001fbdf2b",
		Amount:   10,
		Date:     time.Now(),
	}
	err = kafkaClient.SendTransaction(&transaction)
	if err != nil {
		log.Println("Failed sending transaction %s", err)
	}

}

func basicKafkaListener(stop chan bool) {
	// to consume messages
	topic := "kafka-transaction"
	partition := 0

	conn, _ := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", topic, partition)

	_ = conn.SetReadDeadline(time.Now().Add(10 * time.Second))
	batch := conn.ReadBatch(10e3, 1e6) // fetch 10KB min, 1MB max

	b := make([]byte, 10e3) // 10KB max per message
	for {
		_, err := batch.Read(b)
		if err != nil {
			break
		}
		fmt.Println(string(b))
	}
	stop <- true
	_ = batch.Close()
	_ = conn.Close()
}

func basicPayingClient(originid string, destinationId string) {
	// use the generated client stub
	service, err := grpc.Dial("localhost:50052", grpc.WithInsecure())
	if err != nil {
		println(err)
	}
	cl := transferService.NewTransferValidatorServiceClient(service)

	answer, err := cl.Pay(context.TODO(), &transferService.Transfer{
		Origin:      originid,
		Destination: destinationId,
		Amount:      200,
	})
	if err != nil {
		log.Print(answer)
		println("error", err)
	}
}
