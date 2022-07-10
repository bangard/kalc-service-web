package handlers

import (
	"dobledcloud.com/consumers/server"
	"encoding/json"
	"net/http"
)

type HealthResponse struct {
	Message string `json:"message"`
	Status  bool   `json:"status"`
}

func HealthHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(HealthResponse{
			Message: "Health Check",
			Status:  true,
		})
	}
}
