package app

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"
)

func timeHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	timezonesStr := r.Form.Get("tz")
	timezones := strings.Split(timezonesStr, ",")

	timeMap := make(map[string]string)

	for _, tz := range timezones {
		loc, err := time.LoadLocation(tz)
		if err != nil {
			timeMap[tz] = "Invalid timezone"
		} else {
			currentTime := time.Now().In(loc).Format(time.RFC3339)
			timeMap[tz] = currentTime
		}
	}

	responseJSON, err := json.Marshal(timeMap)
	if err != nil {
		http.Error(w, "Error generating JSON response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(responseJSON)
}
