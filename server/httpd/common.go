package httpd

// contains frequently used helper functions for http

import (
	"encoding/json"
	"net/http"
)

func WriteJSON(w http.ResponseWriter, status int, data any) error {
    w.WriteHeader(status)
	return json.NewEncoder(w).Encode(data)
}
