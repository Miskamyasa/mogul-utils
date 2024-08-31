package response

import (
	"encoding/json"
	"github.com/Miskamyasa/mogul-utils/notify"
	"net/http"
)

func SendJsonResponse(w http.ResponseWriter, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(payload)
	if err != nil {
		notify.Send("Error encoding the response", err)
		return
	}
}

func SendInternalServerError(w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
	_, err := w.Write([]byte("Internal Server Error"))
	if err != nil {
		notify.Send("Error writing the response", err)
	}
}

func SendBadRequest(w http.ResponseWriter) {
	w.WriteHeader(http.StatusBadRequest)
	_, err := w.Write([]byte("Bad Request"))
	if err != nil {
		notify.Send("Error writing the response", err)
	}
}
