package models

type ApiResponse struct {
	Services []struct {
		AtocName       string `json:"atocName"`
		LocationDetail struct {
			GbttBookedArrival   string `json:"gbttBookedArrival"`
			GbttBookedDeparture string `json:"gbttBookedDeparture"`
			RealtimeArrival     string `json:"realtimeArrival"`
			RealtimeDeparture   string `json:"realtimeDeparture"`
			Destination         []struct {
				Description string `json:"description"`
			} `json:"destination"`
		} `json:"locationDetail"`
	} `json:"services"`
}

type ParsedResponse struct {
	Departures []Departures `json:"departures"`
}

type Departures struct {
	Notes              string `json:"notes"`
	AimedDepartureTime string `json:"aimedDepartureTime"`
	DestinationName    string `json:"destinationStation"`
	CallingAt          string `json:"callingAt"`
	Status             string `json:"status"`
}
