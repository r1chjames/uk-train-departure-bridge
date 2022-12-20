package implementation

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"uk-train-departure-lambda/models"
)

func buildApiRequest(config models.AppConfig) *http.Request {
	req, err := http.NewRequest("GET", fmt.Sprintf("https://api.rtt.io/api/v1/json/search/%s/to/%s", config.DepartureStation, config.DestinationStation), nil)
	req.SetBasicAuth(config.Username, config.Password)

	if err != nil {
		log.Fatal(err)
	}

	return req
}

func parseRequestBodyIntoApiResponseStruct(body []byte) models.ApiResponse {

	var apiResponse models.ApiResponse
	if err := json.Unmarshal(body, &apiResponse); err != nil {
		log.Fatalf("Failed marshalling the API response %s\n", err)
	}
	return apiResponse
}

func performApiRequest(req *http.Request) []byte {
	client := &http.Client{}
	resp, err := client.Do(req)
	log.Printf("Making API request to %s\n", req.URL)

	var body []byte
	if err != nil {
		log.Fatalf("The HTTP request failed with error %s\n", err)
	} else {
		body, _ = io.ReadAll(resp.Body)
	}
	return body
}

func QueryUpcomingTrains(config models.AppConfig) (interface{}, error) {

	req := buildApiRequest(config)
	body := performApiRequest(req)
	log.Printf("API response body %s", body)
	entireStruct := parseRequestBodyIntoApiResponseStruct(body)
	return parseFullRequest(config, entireStruct), nil
}

func parseDepartureTime(time string) string {
	return fmt.Sprintf("%s:%s", time[0:2], time[2:4])
}

func parseFullRequest(config models.AppConfig, response models.ApiResponse) models.ParsedResponse {
	var fullResponse models.ParsedResponse
	for i, train := range response.Services {
		if i <= config.FutureTrainCount-1 {
			var status string
			if train.LocationDetail.GbttBookedDeparture == train.LocationDetail.RealtimeDeparture {
				status = "On Time"
			} else {
				status = fmt.Sprintf("Exp %s", parseDepartureTime(train.LocationDetail.RealtimeArrival))
			}

			departure := models.Departures{
				Notes:              fmt.Sprintf("Operated by %s", train.AtocName),
				AimedDepartureTime: parseDepartureTime(train.LocationDetail.GbttBookedDeparture),
				DestinationName:    train.LocationDetail.Destination[0].Description,
				CallingAt:          train.LocationDetail.Destination[0].Description,
				Status:             status,
			}

			fullResponse.Departures = append(fullResponse.Departures, departure)

		}
	}
	return fullResponse
}
