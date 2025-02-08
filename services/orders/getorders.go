package services

import (
	"context"
	"oms/utils/dbconn"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func GetOrdersBySeller(c *gin.Context, sellerID string) []bson.M {
	collection := dbconn.DB_Instance.Database("your_db").Collection("orders")
	filter := bson.M{"seller_id": sellerID}

	var orders []bson.M
	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		c.JSON(500, gin.H{"error": "Unable to fetch orders"})
		return orders
	}
	defer cursor.Close(context.TODO())

	if err := cursor.All(context.TODO(), &orders); err != nil {
		c.JSON(500, gin.H{"error": "Error parsing orders"})
		return orders
	}

	return orders
}
