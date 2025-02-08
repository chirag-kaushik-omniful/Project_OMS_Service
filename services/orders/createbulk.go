package services

import (
	"context"
	"encoding/json"
	"fmt"

	// "oms/utils/csv"

	"oms/utils/csv"
	"oms/utils/intsrv"
	"oms/utils/kafka"

	"github.com/omniful/go_commons/pubsub"
)

type Order struct {
	SNo      string `json:"sno"`
	SellerID string `json:"seller_id"`
	OrderID  string `json:"order_id"`
	ItemID   string `json:"item_id"`
	Quantity string `json:"quantity"`
	Status   string `json:"status"`
}

type OrderResponse struct {
	ValidOrders   []Order `json:"valid_orders"`
	MissingOrders []Order `json:"missing_orders"`
}

func CreateBulkOrder(filePath string) {
	data := csv.ParseCSV(filePath)
	url := "/sku/verify"
	var resp OrderResponse
	intsrv.PostReq(context.Background(), &resp, url, data)
	dd, _ := json.Marshal(resp)
	fmt.Println("DATA: ", string(dd))

	// collection := dbconn.DB_Instance.Database("orders").Collection("orders")

	// for _, item := range resp.ValidOrders {
	// 	order := modals.Order{
	// 		ID:       primitive.NewObjectID(),
	// 		SellerID: item.SellerID,
	// 		// TotalAmount: item.TotalAmount,
	// 		Status:    "Pending",
	// 		CreatedAt: time.Now(),
	// 		UpdatedAt: time.Now(),
	// 		Items: []modals.OrderItem{
	// 			{
	// 				ID:        primitive.NewObjectID(),
	// 				ItemID:    item.ItemID,
	// 				Quantity:  item.Quantity,
	// 				CreatedAt: time.Now(),
	// 				UpdatedAt: time.Now(),
	// 			},
	// 		},
	// 	}

	// 	_, err := collection.InsertOne(context.TODO(), order)
	// 	if err != nil {
	// 		log.Fatalf("Failed to insert order: %v", err)
	// 	}

	// 	fmt.Println("Order inserted successfully")
	// }

	// err := appendToCSV("failed_orders.csv", resp.MissingOrders)
	// if err != nil {
	// 	return
	// }

	fmt.Println("CSV Writed successfully!")

	// Create message with key for FIFO ordering
	msg := &pubsub.Message{
		Topic: "my-topic",
		Key:   "customer-123",
		Value: []byte("Hello Kafka!"),
		Headers: map[string]string{
			"custom-header": "value",
		},
	}

	// Context with request ID
	ctx := context.WithValue(context.Background(), "request_id", "req-123")

	// Synchronous publish - HeaderXOmnifulRequestID will be automatically added
	err := kafka.Producer.Publish(ctx, msg)
	if err != nil {
		panic(err)
	}

}

// func appendToCSV(filePath string, orders []Order) error {
// 	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
// 	if err != nil {
// 		return err
// 	}
// 	defer file.Close()

// 	writer := csv.NewWriter(file)
// 	defer writer.Flush()

// 	for _, order := range orders {
// 		row := []string{order.SNo, order.SellerID, order.OrderID, order.ItemID, order.Quantity, order.Status}
// 		if err := writer.Write(row); err != nil {
// 			return err
// 		}
// 	}
// 	return nil
// }
