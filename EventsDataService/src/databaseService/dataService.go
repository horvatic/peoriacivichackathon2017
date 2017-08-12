package databaseService

import (
	"models"
	"strings"
)

func GetEventData(eventType string) []models.Event {
	if strings.Compare(eventType, "BikeTrail") == 0 {
		model1 := models.Event{Length: "10", Description: "Bike Route South Side", Date: "9/1/2017"}
		model1.Latitude = []string{"-89.618111585971818", "40.685981859253047"}
		model1.Longitude = []string{"-80", "35"}
		model2 := models.Event{Length: "7", Description: "Bike Route North Side", Date: "9/5/2017"}
		model2.Latitude = []string{"-81.6171818", "30.1859253047"}
		model2.Longitude = []string{"-70", "65"}
		return []models.Event{model1, model2}
	} else if strings.Compare(eventType, "Bus") == 0 {
		model := models.Event{Length: "30", Description: "Bus Route", Date: "5/8/2012"}
		model.Latitude = []string{"-79.618111585971818", "50.685981859253047"}
		model.Longitude = []string{"-50", "95"}
		return []models.Event{model}
	}
	return []models.Event{models.Event{}}
}
