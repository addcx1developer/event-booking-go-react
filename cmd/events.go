package main

import (
	"encoding/json"
	"net/http"
)

func healthHandler(w http.ResponseWriter, r *http.Request) {
	var payload = struct {
		Status  string `json:"status"`
		Message string `json:"message"`
		Version string `json:"version"`
	}{
		Status:  "Active",
		Message: "Event Booking API is up and running",
		Version: "1.0.0",
	}

	if err := json.NewEncoder(w).Encode(payload); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
