package services

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	// "oms/utils/csv"

	"oms/repositories"
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

	repositories.CreateBulkOrderInDB(resp.ValidOrders)

	err := appendToCSV("failed_orders.csv", resp.MissingOrders)
	if err != nil {
		return
	}

	fmt.Println("CSV Writed successfully!")

}

func appendToCSV(filePath string, orders []Order) error {
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, order := range orders {
		row := []string{order.SNo, order.SellerID, order.OrderID, order.ItemID, order.Quantity, order.Status}
		if err := writer.Write(row); err != nil {
			return err
		}
	}
	return nil
}
