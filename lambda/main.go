package main

import (
	"uk-train-departure-lambda/implementation"
	"uk-train-departure-lambda/models"
	"uk-train-departure-lambda/utils"
)

func main() {
	appConfig := parseAppVariables()
	upcomingTrains, _ := implementation.QueryUpcomingTrains(appConfig)
	implementation.PublishMessageToIot(appConfig, upcomingTrains.(models.ParsedResponse))
}

func parseAppVariables() models.AppConfig {
	return models.AppConfig{
		DepartureStation:   utils.GetStringEnv("DEPARTURE_STATION", "KGX"),
		DestinationStation: utils.GetStringEnv("DESTINATION_STATION", "LDS"),
		FutureTrainCount: utils.GetIntEnv("FUTURE_TRAIN_COUNT", 3),
		RefreshIntervalSec: utils.GetIntEnv("REFRESH_INTERVAL_SEC", 300),
		OperatingHours:     utils.GetStringEnv("OPERATING_HOURS", "06_22"),
		IotTopic: utils.GetStringEnv("IOT_TOPIC", "LDS"),
		TransportApiUrl: "https://transportapi.com/v3/uk/places.json",
		ApiKey:             utils.GetStringEnv("API_KEY", ""),
		AppId:              utils.GetStringEnv("APP_ID", ""),
	}
}


