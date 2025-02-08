package services

import (
	"context"
	"encoding/json"
	"fmt"

	// "oms/utils/csv"

	"oms/utils/csv"
	"oms/utils/intsrv"
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

}
