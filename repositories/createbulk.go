package repositories

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"oms/modals"
	"oms/utils/dbconn"
	"oms/utils/kafka"
	"time"

	"github.com/omniful/go_commons/pubsub"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Order struct {
	SNo      string `json:"sno"`
	SellerID string `json:"seller_id"`
	OrderID  string `json:"order_id"`
	ItemID   string `json:"item_id"`
	Quantity string `json:"quantity"`
	Status   string `json:"status"`
}

type Req struct {
	OrderID  string `json:"order_id"`
	ItemID   string `json:"item_id"`
	Quantity string `json:"quantity"`
}

func CreateBulkOrderInDB(ValidOrders []Order) {

	collection := dbconn.DB_Instance.Database("oms_db").Collection("orders")
	for _, item := range ValidOrders {
		order := modals.Order{
			ID:       primitive.NewObjectID(),
			SellerID: item.SellerID,
			// TotalAmount: item.TotalAmount,
			Status:    item.Status,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Items: []modals.OrderItem{
				{
					ID:        primitive.NewObjectID(),
					ItemID:    item.ItemID,
					Quantity:  item.Quantity,
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				},
			},
		}

		_, err := collection.InsertOne(context.TODO(), order)
		if err != nil {
			log.Fatalf("Failed to insert order: %v", err)
		}

		// Create message with key for FIFO ordering
		req := &Req{
			OrderID:  item.OrderID,
			ItemID:   item.ItemID,
			Quantity: item.Quantity,
		}
		jsonData, _ := json.Marshal(req)
		msg := &pubsub.Message{
			Topic: "my-topic",
			Key:   "customer-123",
			Value: jsonData,
			Headers: map[string]string{
				"custom-header": "value",
			},
		}

		// Context with request ID
		ctx := context.WithValue(context.Background(), "request_id", "req-123")

		err = kafka.Producer.Publish(ctx, msg)
		if err != nil {
			panic(err)
		}

	}
	fmt.Println("Order inserted successfully")

}
