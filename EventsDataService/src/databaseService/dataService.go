package databaseService

import (
	"models"
	"strings"
)

func GetEventData(eventType string) models.Event {
	if strings.Compare(eventType, "BikeTrail") == 0 {
		model := models.Event{Length: "10", Description: "Bike Route South Side", Date: "9/1/2017"}
		model.Latitude = []string{"-89.618111585971818", "40.685981859253047"}
		model.Longitude = []string{"-80", "35"}
		return model
	} else if strings.Compare(eventType, "Bus") == 0 {
		model := models.Event{Length: "30", Description: "Bus Route", Date: "5/8/2012"}
		model.Latitude = []string{"-79.618111585971818", "50.685981859253047"}
		model.Longitude = []string{"-50", "95"}
		return model
	}
	return models.Event{}

}
