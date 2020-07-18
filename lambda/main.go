package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"uk-train-departure-lambda/implementation"
	"uk-train-departure-lambda/models"
	"uk-train-departure-lambda/utils"
)

func main() {
	lambda.Start(fetchTrainTimes)
}

func fetchTrainTimes() (string, error) {
	appConfig := parseAppVariables()
	upcomingTrains, _ := implementation.QueryUpcomingTrains(appConfig)
	err := implementation.PublishMessageToIot(appConfig, upcomingTrains.(models.ParsedResponse))
	return "Completed", err
}

func parseAppVariables() models.AppConfig {
	return models.AppConfig{
		DepartureStation:   utils.GetStringEnv("DEPARTURE_STATION", "KGX"),
		DestinationStation: utils.GetStringEnv("DESTINATION_STATION", "LDS"),
		FutureTrainCount: utils.GetIntEnv("FUTURE_TRAIN_COUNT", 3),
		OperatingHours:     utils.GetStringEnv("OPERATING_HOURS", "06_22"),
		IotTopic: utils.GetStringEnv("IOT_TOPIC", "LDS"),
		TransportApiUrl: "https://transportapi.com/v3/uk/places.json",
		ApiKey:             utils.GetStringEnv("API_KEY", ""),
		AppId:              utils.GetStringEnv("APP_ID", ""),
	}
}


