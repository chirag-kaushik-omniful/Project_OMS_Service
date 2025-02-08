package kafka

import (
	"context"
	"time"

	"github.com/omniful/go_commons/kafka"
	"github.com/omniful/go_commons/pubsub"
	"github.com/omniful/go_commons/pubsub/interceptor"
)

// Implement message handler
type MessageHandler struct{}

func (h *MessageHandler) Process(ctx context.Context, msg *pubsub.Message) error {
	// Process message
	return nil
}

func InitKafka() {
	// Initialize producer with configuration
	producer := kafka.NewProducer(
		kafka.WithBrokers([]string{"localhost:9092"}),
		kafka.WithClientID("my-producer"),
		kafka.WithKafkaVersion("2.8.1"),
	)
	defer producer.Close()

	// Create message with key for FIFO ordering
	msg := &pubsub.Message{
		Topic: "my-topic",
		// Key is crucial for maintaining FIFO ordering
		// Messages with the same key will be delivered to the same partition in order
		Key:   "customer-123",
		Value: []byte("Hello Kafka!"),
		Headers: map[string]string{
			"custom-header": "value",
			// Note: HeaderXOmnifulRequestID will be automatically added
			// from context if present
		},
	}

	// Context with request ID
	ctx := context.WithValue(context.Background(), "request_id", "req-123")

	// Synchronous publish - HeaderXOmnifulRequestID will be automatically added
	err := producer.Publish(ctx, msg)
	if err != nil {
		panic(err)
	}

	// Batch publish with consistent keys for ordering
	messages := []*pubsub.Message{
		{
			Topic: "my-topic",
			Key:   "customer-123", // Same key maintains ordering
			Value: []byte("Message 1"),
		},
		{
			Topic: "my-topic",
			Key:   "customer-123", // Same key maintains ordering
			Value: []byte("Message 2"),
		},
	}
	err = producer.PublishBatch(ctx, messages)
	if err != nil {
		panic(err)
	}

	// Initialize consumer with configuration
	consumer := kafka.NewConsumer(
		kafka.WithBrokers([]string{"localhost:9092"}),
		kafka.WithConsumerGroup("my-consumer-group"),
		kafka.WithClientID("my-consumer"),
		kafka.WithKafkaVersion("2.8.1"),
		kafka.WithRetryInterval(time.Second),
		kafka.WithDeadLetterConfig(&kafka.DeadLetterQueueConfig{
			Queue:     "dlq-queue",
			Account:   "aws-account",
			Region:    "us-east-1",
			ShouldLog: true,
		}),
	)
	defer consumer.Close()

	// Set NewRelic interceptor for monitoring
	consumer.SetInterceptor(interceptor.NewRelicInterceptor())

	// Register message handler for topic
	handler := &MessageHandler{}
	consumer.RegisterHandler("my-topic", handler)

	// Start consuming messages
	ctx = context.Background()
	consumer.Subscribe(ctx)

}
