package response

import (
	"encoding/json"
	"net/http"

	"github.com/Miskamyasa/utils/alerts"
)

func SendJsonResponse(w http.ResponseWriter, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(payload)
	if err != nil {
		alerts.Send("Error encoding the response", err)
		return
	}
}

func SendInternalServerError(w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
	_, err := w.Write([]byte("Internal Server Error"))
	if err != nil {
		alerts.Send("Error writing the response", err)
	}
}

func SendBadRequest(w http.ResponseWriter, msg string) {
	w.WriteHeader(http.StatusBadRequest)
	_, err := w.Write([]byte("Bad Request! " + msg))
	if err != nil {
		alerts.Send("Error writing the response", err)
	}
}

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte("OK"))
	if err != nil {
		alerts.Send("Error writing the response", err)
		return
	}
}
