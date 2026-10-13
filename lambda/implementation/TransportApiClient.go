package implementation

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"time"
	"uk-train-departure-lambda/models"
)

const apiBaseUrl = "https://data.rtt.io"

// Wide enough to catch the next few departures on hourly routes.
const searchWindowMinutes = 120

func buildApiRequest(config models.AppConfig, accessToken string) *http.Request {
	query := url.Values{
		"code":       {config.DepartureStation},
		"filterTo":   {config.DestinationStation},
		"timeWindow": {strconv.Itoa(searchWindowMinutes)},
	}
	req, err := http.NewRequest("GET", apiBaseUrl+"/gb-nr/location?"+query.Encode(), nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Authorization", "Bearer "+accessToken)
	return req
}

// The portal issues a refresh token, which has to be swapped for a short life access token.
func acquireAccessToken(config models.AppConfig) (string, error) {
	req, err := http.NewRequest("GET", apiBaseUrl+"/api/get_access_token", nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("Authorization", "Bearer "+config.Token)

	body, err := performApiRequest(req)
	if err != nil {
		return "", err
	}

	var tokenResponse models.AccessTokenResponse
	if err := json.Unmarshal(body, &tokenResponse); err != nil {
		return "", fmt.Errorf("failed marshalling the access token response %s: %w", body, err)
	}
	return tokenResponse.Token, nil
}

func parseRequestBodyIntoApiResponseStruct(body []byte) (models.ApiResponse, error) {
	var apiResponse models.ApiResponse
	if err := json.Unmarshal(body, &apiResponse); err != nil {
		return apiResponse, fmt.Errorf("failed marshalling the API response %s: %w", body, err)
	}
	return apiResponse, nil
}

func performApiRequest(req *http.Request) ([]byte, error) {
	client := &http.Client{}
	log.Printf("Making API request to %s\n", req.URL)

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("the HTTP request to %s failed: %w", req.URL, err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	if resp.StatusCode == http.StatusNoContent {
		return nil, nil // 204: valid query, no services in the window
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API returned %s: %s", resp.Status, body)
	}
	return body, nil
}

func QueryUpcomingTrains(config models.AppConfig) (models.ParsedResponse, error) {

	accessToken, err := acquireAccessToken(config)
	if err != nil {
		return models.ParsedResponse{}, err
	}

	req := buildApiRequest(config, accessToken)
	body, err := performApiRequest(req)
	if err != nil {
		return models.ParsedResponse{}, err
	}
	if len(body) == 0 {
		return models.ParsedResponse{}, nil
	}

	log.Printf("API response body %s", body)
	entireStruct, err := parseRequestBodyIntoApiResponseStruct(body)
	if err != nil {
		return models.ParsedResponse{}, err
	}
	return parseFullRequest(config, entireStruct), nil
}

// The previous API returned local HHMM times, the new one returns ISO 8601 datetimes.
func parseDepartureTime(datetime string) string {
	for _, layout := range []string{time.RFC3339, "2006-01-02T15:04:05"} {
		if parsed, err := time.Parse(layout, datetime); err == nil {
			return parsed.Format("15:04")
		}
	}
	return datetime
}

func firstNonEmpty(values ...string) string {
	for _, value := range values {
		if value != "" {
			return value
		}
	}
	return ""
}

func parseFullRequest(config models.AppConfig, response models.ApiResponse) models.ParsedResponse {
	var fullResponse models.ParsedResponse
	for i, train := range response.Services {
		if i >= config.FutureTrainCount {
			break
		}

		departure := train.TemporalData.Departure
		expected := firstNonEmpty(departure.RealtimeActual, departure.RealtimeForecast, departure.RealtimeEstimate)

		status := "On Time"
		if expected != "" && expected != departure.ScheduleAdvertised {
			status = fmt.Sprintf("Exp %s", parseDepartureTime(expected))
		}

		destination := ""
		if len(train.Destination) > 0 {
			destination = train.Destination[0].Location.Description
		}

		fullResponse.Departures = append(fullResponse.Departures, models.Departures{
			Notes:              fmt.Sprintf("Operated by %s", train.ScheduleMetadata.Operator.Name),
			AimedDepartureTime: parseDepartureTime(departure.ScheduleAdvertised),
			DestinationName:    destination,
			CallingAt:          destination,
			Status:             status,
		})
	}
	return fullResponse
}
