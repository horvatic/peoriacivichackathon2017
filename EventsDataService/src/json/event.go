package json

import (
	"databaseService"
	"encoding/json"
	"net/http"
	"net/url"
)

func EventData(res http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()
	res.Header().Set("Content-Type", "application/json")
	message := databaseService.GetEventData(getEvent(req.URL))
	json.NewEncoder(res).Encode(message)
}

func getEvent(u *url.URL) string {
	q := u.Query()
	return q.Get("event")
}
