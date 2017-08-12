package models

type Event struct {
	Length      string   `json: "length"`
	Latitude    []string `json "latitude`
	Longitude   []string `json "longitude"`
	Description string   `json "description"`
	Date        string   `json "date"`
}
