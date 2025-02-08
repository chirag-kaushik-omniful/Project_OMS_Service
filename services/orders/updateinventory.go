package services

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"oms/modals"
	"oms/utils/dbconn"
	"oms/utils/intsrv"

	"go.mongodb.org/mongo-driver/bson"
)

type Response struct {
	Message string `json:"message"`
	Status  string `json:"status"`
}

func UpdateInventory(bytedata []byte) {
	var data modals.Order

	json.Unmarshal(bytedata, &data)

	url := "/inventory/edit"
	var resp Response
	intsrv.PostReq(context.Background(), &resp, url, data)
	dd, _ := json.Marshal(resp)
	fmt.Println("DATA: ", string(dd))

	if resp.Status == string(http.StatusOK) {
		// update order status
		collection := dbconn.DB_Instance.Database("oms_db").Collection("orders")
		filter := bson.M{"OrderID": data.OrderID}
		update := bson.M{"$set": bson.M{"status": "new_order"}}

		_, err := collection.UpdateOne(context.TODO(), filter, update)
		if err != nil {
			log.Fatal(err)
		}
	}
}
