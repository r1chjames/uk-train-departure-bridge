package utils

import "uk-train-departure-lambda/models"

func FindStationIndex(all models.DestinationList, stationToFind string) int {
	for i, station := range all.Stops {
		if station.StationCode == stationToFind {
			return i
		}
	}
	return 0
}
