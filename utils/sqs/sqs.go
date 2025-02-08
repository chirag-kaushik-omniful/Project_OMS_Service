package sqs

import (
	"context"
	"fmt"
	config "oms/configs"
	services "oms/services/orders"

	"github.com/omniful/go_commons/log"
	"github.com/omniful/go_commons/sqs"
)

var Queue *sqs.Queue
var Consumer *sqs.Consumer
var Publisher *sqs.Publisher

type MyHandler struct{}

func (h *MyHandler) Handle(msg sqs.Message) error {
	fmt.Println("Processing message:", string(msg.Value))
	services.CreateBulkOrder(string(msg.Value))
	return nil
}

func (h *MyHandler) Process(ctx context.Context, msgs *[]sqs.Message) error {
	for _, msg := range *msgs {
		if err := h.Handle(msg); err != nil {
			return err
		}
	}
	return nil
}

func InitSQS() {
	queueName := "OMSQueue.fifo"
	queue, err := sqs.NewFifoQueue(context.Background(), queueName, &sqs.Config{
		Account:  config.SQS_Config.Account,
		Endpoint: config.SQS_Config.Endpoint,
		Region:   config.SQS_Config.Region,
	})
	if err != nil || queue == nil {
		log.Errorf("initialization error. queue: %v, err : %v, publisher: %+v", queueName, err, queue)
		return
	}
	Queue = queue

	Handler := &MyHandler{}
	log.Debugf("creating queue: %v", queueName)
	consumer, err := sqs.NewConsumer(
		queue,
		uint64(1),
		4,
		Handler,
		10,
		30,
		true,
		false,
	)
	if err != nil || consumer == nil {
		log.Errorf("initialization error. queue: %v, err : %v, publisher: %+v", queueName, err, consumer)
		return
	}
	Consumer = consumer

	consumer.Start(context.Background())
	fmt.Println("queue created successfully")

	publisher := sqs.NewPublisher(queue)
	Publisher = publisher

}
