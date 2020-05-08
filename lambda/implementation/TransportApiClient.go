package implementation

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"uk-train-departure-lambda/models"
)

func buildApiRequest(config models.AppConfig) string {
	req, err := http.NewRequest("GET", fmt.Sprintf("https://transportapi.com/v3/uk/train/station/%s/live.json", config.DepartureStation), nil)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}

	query := req.URL.Query()
	query.Add("calling_at", config.DestinationStation)
	query.Add("app_id", config.AppId)
	query.Add("app_key", config.ApiKey)
	req.URL.RawQuery = query.Encode()

	return req.URL.String()
}

func parseRequestBodyIntoApiResponseStruct(body []byte) models.ApiResponse {

	var apiResponse models.ApiResponse
	if err := json.Unmarshal(body, &apiResponse); err != nil {
		fmt.Printf("Failed marshalling the API response %s\n", err)
	}
	return apiResponse
}

func parseRequestBodyIntoApiDestinationResponseStruct(body []byte) models.DestinationList {

	var apiResponse models.DestinationList
	if err := json.Unmarshal(body, &apiResponse); err != nil {
		fmt.Printf("Failed marshalling the API response %s\n", err)
	}
	return apiResponse
}

func performApiRequest(apiUrl string) []byte {
	fmt.Println(apiUrl)
	resp, err := http.Get(apiUrl)

	var body []byte
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		body, _ = ioutil.ReadAll(resp.Body)
	}
	return body
}

func QueryUpcomingTrains(config models.AppConfig) (interface{}, error) {

	apiUrl := buildApiRequest(config)
	body := performApiRequest(apiUrl)
	entireStruct := parseRequestBodyIntoApiResponseStruct(body)
	return parseFullRequest(config, entireStruct), nil
}

func parseDestinationList(destinationsAll models.DestinationList) string {
	var stationNames []string
	for _, station := range destinationsAll.Stops {
		stationNames = append(stationNames, station.StationName)
	}
	return strings.Join(stationNames, ", ")
}

func parseFullRequest(config models.AppConfig, response models.ApiResponse) models.ParsedResponse {
	var fullResponse models.ParsedResponse
	for i, train := range response.Departures.All {
		if i <= config.FutureTrainCount-1 {
			var status string
			if train.AimedDepartureTime == train.ExpectedDepartureTime {
				status = "On Time"
			} else {
				status = train.ExpectedDepartureTime
			}

			destinationsResponse := performApiRequest(train.ServiceTimetable.ID)
			parsedDestinationList := parseRequestBodyIntoApiDestinationResponseStruct(destinationsResponse)
			callingStations := parseDestinationList(parsedDestinationList)


			departure := models.Departures{
				OperatorName:       train.Operator,
				AimedDepartureTime: train.AimedDepartureTime,
				DestinationName:    train.DestinationName,
				CallingAt:   		callingStations,
				Status:             status,
			}

			fullResponse.Departures = append(fullResponse.Departures, departure)

		}
	}
	return fullResponse
}