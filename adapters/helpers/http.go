package helpers

import (
	"encoding/json"
	"net/http"
)

// SuccessReponse makes the response with payload as json format
func SuccessReponse(w http.ResponseWriter, payload interface{}, status int) {
	response, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(status)
	w.Write([]byte(response))
}

// ErrorResponse makes an error response with message and status code
func ErrorResponse(w http.ResponseWriter, errorMessage string, code int) {
	SuccessReponse(w, map[string]string{"error": errorMessage}, code)
}
