package controllers

import (
	"context"
	"net/http"
	services "oms/services/orders"
	"oms/utils/sqs"

	"github.com/gin-gonic/gin"
	"github.com/omniful/go_commons/log"
	SQS "github.com/omniful/go_commons/sqs"
)

func ViewOrders(c *gin.Context) {
	sellerID := c.Query("sellerID")
	if sellerID == "" {
		c.JSON(400, gin.H{"error": "sellerID is required"})
		return
	}

	orders := services.GetOrdersBySeller(c, sellerID)

	c.JSON(http.StatusOK, gin.H{"orders": orders})
}

func CreateBulkOrder(c *gin.Context) {
	var req struct {
		Address string `json:"address"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	message := &SQS.Message{
		GroupId:         "group-1",
		Value:           []byte(req.Address),
		ReceiptHandle:   "group-1",
		DeduplicationId: "gp-1",
	}

	ctx := context.Background()
	if err := sqs.Publisher.Publish(ctx, message); err != nil {
		log.Errorf("Failed to publish message: %v", err)
	}

	// services.CreateBulkOrder(req.Address)

	c.JSON(http.StatusOK, gin.H{"message": "success"})
}
