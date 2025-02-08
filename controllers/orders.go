package controllers

import (
	services "oms/services/orders"

	"github.com/gin-gonic/gin"
)

func ViewOrders(c *gin.Context) {
	sellerID := c.Query("sellerID")
	if sellerID == "" {
		c.JSON(400, gin.H{"error": "sellerID is required"})
		return
	}

	// orders := getOrdersBySeller(sellerID)

	c.JSON(200, gin.H{"orders": "orders"})
}

func CreateBulkOrder(c *gin.Context) {
	var req struct {
		Address string `json:"address"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	// message := &SQS.Message{
	// 	GroupId:         "group-1",
	// 	Value:           []byte(req.Address),
	// 	ReceiptHandle:   "group-1",
	// 	DeduplicationId: "gp-1",
	// }

	// ctx := context.Background()
	// if err := sqs.Publisher.Publish(ctx, message); err != nil {
	// 	log.Errorf("Failed to publish message: %v", err)
	// }

	services.CreateBulkOrder(req.Address)

	c.JSON(200, gin.H{"data": "data"})
}
