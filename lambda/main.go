package main

import (
	"log"
	"uk-train-departure-lambda/implementation"
	"uk-train-departure-lambda/models"
	"uk-train-departure-lambda/utils"

	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	log.Printf("Running...")
	lambda.Start(fetchTrainTimes)
}

func fetchTrainTimes() (string, error) {
	appConfig := parseAppVariables()
	upcomingTrains, err := implementation.QueryUpcomingTrains(appConfig)
	if err != nil {
		return "", err
	}
	return "Completed", implementation.PublishMessageToIot(appConfig, upcomingTrains)
}

func parseAppVariables() models.AppConfig {
	return models.AppConfig{
		DepartureStation:   utils.GetStringEnv("DEPARTURE_STATION", "WDS"),
		DestinationStation: utils.GetStringEnv("DESTINATION_STATION", "LDS"),
		FutureTrainCount:   utils.GetIntEnv("FUTURE_TRAIN_COUNT", 3),
		IotTopic:           utils.GetStringEnv("IOT_TOPIC", "LDS"),
		Token:              utils.GetStringEnv("API_TOKEN", ""),
	}
}
