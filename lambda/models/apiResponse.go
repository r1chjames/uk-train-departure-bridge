package models

type ApiResponse struct {
	Services []Service `json:"services"`
}

type Service struct {
	ScheduleMetadata struct {
		Operator struct {
			Name string `json:"name"`
		} `json:"operator"`
	} `json:"scheduleMetadata"`
	TemporalData struct {
		Departure TemporalData `json:"departure"`
	} `json:"temporalData"`
	Destination []LocationPair `json:"destination"`
}

type TemporalData struct {
	ScheduleAdvertised string `json:"scheduleAdvertised"`
	RealtimeForecast   string `json:"realtimeForecast"`
	RealtimeEstimate   string `json:"realtimeEstimate"`
	RealtimeActual     string `json:"realtimeActual"`
}

type LocationPair struct {
	Location struct {
		Description string `json:"description"`
	} `json:"location"`
}

type AccessTokenResponse struct {
	Token      string `json:"token"`
	ValidUntil string `json:"validUntil"`
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
