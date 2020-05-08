package models

import "time"

type ApiResponse struct {
	Date        string    `json:"date"`
	TimeOfDay   string    `json:"time_of_day"`
	RequestTime time.Time `json:"request_time"`
	StationName string    `json:"station_name"`
	StationCode string    `json:"station_code"`
	Departures  struct {
		All []struct {
			Mode               string      `json:"mode"`
			Service            string      `json:"service"`
			TrainUID           string      `json:"train_uid"`
			Platform           string      `json:"platform"`
			Operator           string      `json:"operator"`
			OperatorName       string      `json:"operator_name"`
			AimedDepartureTime string      `json:"aimed_departure_time"`
			AimedArrivalTime   interface{} `json:"aimed_arrival_time"`
			AimedPassTime      interface{} `json:"aimed_pass_time"`
			OriginName         string      `json:"origin_name"`
			DestinationName    string      `json:"destination_name"`
			Source             string      `json:"source"`
			Category           string      `json:"category"`
			ServiceTimetable   struct {
				ID string `json:"id"`
			} `json:"service_timetable"`
			Status                    string      `json:"status"`
			ExpectedArrivalTime       interface{} `json:"expected_arrival_time"`
			ExpectedDepartureTime     string      `json:"expected_departure_time"`
			BestArrivalEstimateMins   interface{} `json:"best_arrival_estimate_mins"`
			BestDepartureEstimateMins int         `json:"best_departure_estimate_mins"`
		} `json:"all"`
	} `json:"departures"`
}

type DestinationList struct {
	Service  string `json:"service"`
	TrainUID string `json:"train_uid"`
	Headcode string `json:"headcode"`
	Toc      struct {
		AtocCode string `json:"atoc_code"`
	} `json:"toc"`
	TrainStatus     string      `json:"train_status"`
	OriginName      string      `json:"origin_name"`
	DestinationName string      `json:"destination_name"`
	StopOfInterest  interface{} `json:"stop_of_interest"`
	Date            string      `json:"date"`
	TimeOfDay       interface{} `json:"time_of_day"`
	Mode            string      `json:"mode"`
	RequestTime     time.Time   `json:"request_time"`
	Category        string      `json:"category"`
	Operator        string      `json:"operator"`
	OperatorName    string      `json:"operator_name"`
	Stops           []struct {
		StationCode           string      `json:"station_code"`
		TiplocCode            string      `json:"tiploc_code"`
		StationName           string      `json:"station_name"`
		StopType              string      `json:"stop_type"`
		Platform              string      `json:"platform"`
		AimedDepartureDate    string      `json:"aimed_departure_date"`
		AimedDepartureTime    string      `json:"aimed_departure_time"`
		AimedArrivalDate      interface{} `json:"aimed_arrival_date"`
		AimedArrivalTime      interface{} `json:"aimed_arrival_time"`
		AimedPassDate         interface{} `json:"aimed_pass_date"`
		AimedPassTime         interface{} `json:"aimed_pass_time"`
		ExpectedDepartureDate interface{} `json:"expected_departure_date"`
		ExpectedDepartureTime interface{} `json:"expected_departure_time"`
		ExpectedArrivalDate   interface{} `json:"expected_arrival_date"`
		ExpectedArrivalTime   interface{} `json:"expected_arrival_time"`
		ExpectedPassDate      interface{} `json:"expected_pass_date"`
		ExpectedPassTime      interface{} `json:"expected_pass_time"`
		Status                interface{} `json:"status"`
	} `json:"stops"`
}

type ParsedResponse struct {
	Departures []Departures `json:"departures"`
}

type Departures struct {
	OperatorName       	string `json:"operator"`
	AimedDepartureTime 	string `json:"aimedDepartureTime"`
	DestinationName    	string `json:"destinationStation"`
	CallingAt	   		string `json:"callingAt"`
	Status				string `json:"status"`
}
