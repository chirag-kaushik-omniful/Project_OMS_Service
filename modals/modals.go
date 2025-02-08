package modals

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Order represents an order document in MongoDB
type Order struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	OrderID     string             `bson:"order_id,omitempty" json:"order_id"`
	SellerID    string             `bson:"seller_id,omitempty" json:"seller_id"`
	Items       []OrderItem        `bson:"items" json:"items"` // Embedded order items
	TotalAmount float64            `bson:"total_amount" json:"total_amount"`
	Status      string             `bson:"status" json:"status"` // e.g., pending, shipped, delivered
	CreatedAt   time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt   time.Time          `bson:"updated_at" json:"updated_at"`
}

// OrderItem is embedded within the Order document
type OrderItem struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	ItemID    string             `bson:"item_id,omitempty" json:"item_id"`
	Quantity  string             `bson:"quantity" json:"quantity"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at" json:"updated_at"`
}
