package services

import (
	"context"
	"encoding/json"
	"github.com/segmentio/kafka-go"
)

type KafkaClient struct {
	writer *kafka.Writer
}

func (kafkaClient *KafkaClient) SendTransaction(transaction *TransactionDTO) error {
	bytes, e := json.Marshal(transaction)
	if (e != nil) {
		return e;
	}
	err := kafkaClient.writer.WriteMessages(context.Background(), kafka.Message{Value: bytes,})
	return err;
}

func InitializeKafkaClient(topic string, kafkaBroker string) (*KafkaClient, error) {
	return &KafkaClient{writer: kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{kafkaBroker},
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	}),
	}, nil
}
