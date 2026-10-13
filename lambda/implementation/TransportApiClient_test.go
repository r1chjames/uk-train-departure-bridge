package implementation

import (
	"os"
	"testing"
	"uk-train-departure-lambda/models"
)

// Real response from /gb-nr/location?code=WDS&filterTo=LDS&timeWindow=120
func loadSampleResponse(t *testing.T) models.ApiResponse {
	t.Helper()
	body, err := os.ReadFile("testdata/location_WDS_LDS.json")
	if err != nil {
		t.Fatalf("unable to read test data: %s", err)
	}
	apiResponse, err := parseRequestBodyIntoApiResponseStruct(body)
	if err != nil {
		t.Fatalf("unexpected parse error: %s", err)
	}
	return apiResponse
}

func TestParseFullRequest(t *testing.T) {
	parsed := parseFullRequest(models.AppConfig{FutureTrainCount: 2}, loadSampleResponse(t))

	if len(parsed.Departures) != 2 {
		t.Fatalf("expected 2 departures, got %d", len(parsed.Departures))
	}

	first := parsed.Departures[0]
	if first.AimedDepartureTime != "20:12" || first.Status != "Exp 20:16" {
		t.Errorf("unexpected first departure: %+v", first)
	}
	if first.Notes != "Operated by Northern" || first.DestinationName != "Leeds" || first.CallingAt != "Leeds" {
		t.Errorf("unexpected first departure: %+v", first)
	}

	second := parsed.Departures[1]
	if second.AimedDepartureTime != "20:43" || second.Status != "On Time" {
		t.Errorf("unexpected second departure: %+v", second)
	}
}

func TestParseDepartureTime(t *testing.T) {
	for input, expected := range map[string]string{
		"2026-10-13T20:12:00":       "20:12", // local time, no offset
		"2026-10-13T20:12:00+01:00": "20:12",
		"2026-01-13T20:12:00Z":      "20:12",
		"2012":                      "2012",
	} {
		if actual := parseDepartureTime(input); actual != expected {
			t.Errorf("parseDepartureTime(%s) = %s, expected %s", input, actual, expected)
		}
	}
}
