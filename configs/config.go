package config

import (
	"github.com/omniful/go_commons/sqs"
)

var SQS_Config = &sqs.Config{
	Account:  "273354671146",
	Region:   "us-east-1",
	Endpoint: "https://sqs.us-east-1.amazonaws.com/273354671146/OMSQueue.fifo",
}
