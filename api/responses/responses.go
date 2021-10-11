package responses

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func SUCCESS(w http.ResponseWriter, statusCode int, data interface{}) {
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		fmt.Fprintf(w, "Error formatting response: %v", err)
	}
}

func ERROR(w http.ResponseWriter, statusCode int, err error) {
	if err != nil {
		SUCCESS(w, statusCode, map[string]string{"error": err.Error()})
		return
	}

	SUCCESS(w, http.StatusBadRequest, nil)
}
