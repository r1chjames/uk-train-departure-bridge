package models

type AppConfig struct {
	DepartureStation string
	DestinationStation string
	FutureTrainCount int
	RefreshIntervalSec int
	OperatingHours string
	IotTopic string
	TransportApiUrl string
	ApiKey string
	AppId string
}