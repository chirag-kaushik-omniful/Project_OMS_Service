package csv

import (
	"context"
	"fmt"
	"log"

	"github.com/omniful/go_commons/csv"
)

type Data []map[string]string

func ParseCSV(filePath string) []map[string]string {
	csvReader, err := csv.NewCommonCSV(
		csv.WithBatchSize(1000),
		csv.WithSource(csv.Local),
		csv.WithLocalFileInfo(filePath),
	)
	if err != nil {
		log.Fatal(err)
	}
	err = csvReader.InitializeReader(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	var data Data
	for !csvReader.IsEOF() {
		var records csv.Records
		records, err = csvReader.ReadNextBatch()
		if err != nil {
			log.Fatal(err)
		}
		// Process the records
		var batchData Data
		var headers csv.Headers = []string{"sno", "seller_id", "order_id", "item_id", "quantity", "status"}
		records.Unmarshal(headers, &batchData)
		fmt.Println(batchData)
		data = append(data, batchData...)
	}

	return data

}
