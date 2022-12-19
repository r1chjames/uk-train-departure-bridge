package implementation

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/iotdataplane"
	"log"
	"uk-train-departure-lambda/models"
)

func createIotSession() *iotdataplane.IoTDataPlane {
	awsSession, _ := session.NewSession()
	return iotdataplane.New(awsSession, &aws.Config{
		Region:   aws.String("eu-west-1"),
		Endpoint: aws.String("https://a1jht9wg8u2etb-ats.iot.eu-west-1.amazonaws.com"),
	})
}

func PublishMessageToIot(appConfig models.AppConfig, message models.ParsedResponse) error {
	reqBodyBytes := new(bytes.Buffer)
	json.NewEncoder(reqBodyBytes).Encode(message)
	fmt.Println(reqBodyBytes)
	payload := reqBodyBytes.Bytes()
	qos := int64(1)

	publishInput := iotdataplane.PublishInput{
		Payload: payload,
		Qos:     &qos,
		Topic:   &appConfig.IotTopic,
		Retain:  aws.Bool(true),
	}

	client := createIotSession()

	_, err := client.Publish(&publishInput)
	if err != nil {
		log.Fatal("Unable to send IoT message", err)
	}
	return err
}
