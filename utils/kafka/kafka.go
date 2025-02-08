package kafka

import (
	"context"
	"fmt"
	"time"

	"github.com/omniful/go_commons/kafka"
	"github.com/omniful/go_commons/pubsub"
	"github.com/omniful/go_commons/pubsub/interceptor"
)

var Producer *kafka.ProducerClient
var Consuemr *kafka.ConsumerClient

// Implement message handler
type MessageHandler struct{}

func (h *MessageHandler) Process(ctx context.Context, msg *pubsub.Message) error {
	// Process message
	fmt.Println(string(msg.Value))
	return nil
}

func InitKafka() {
	// Initialize producer with configuration
	producer := kafka.NewProducer(
		kafka.WithBrokers([]string{"localhost:9092"}),
		kafka.WithClientID("my-producer"),
		kafka.WithKafkaVersion("2.8.1"),
	)

	Producer = producer

	// Initialize consumer with configuration
	consumer := kafka.NewConsumer(
		kafka.WithBrokers([]string{"localhost:9092"}),
		kafka.WithConsumerGroup("my-consumer-group"),
		kafka.WithClientID("my-consumer"),
		kafka.WithKafkaVersion("2.8.1"),
		kafka.WithRetryInterval(time.Second),
	)

	// Set NewRelic interceptor for monitoring
	consumer.SetInterceptor(interceptor.NewRelicInterceptor())

	// Register message handler for topic
	handler := &MessageHandler{}
	consumer.RegisterHandler("my-topic", handler)

	// Start consuming messages
	ctx := context.Background()
	consumer.Subscribe(ctx)

}
