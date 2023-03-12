package middleware

import (
	"encoding/json"
	"net/http"
)

// ErrorMessage structure that returns errors
type ErrorMessage struct {
	Status  int    `json:"status"`
	Error   string `json:"error"`
	Message string `json:"message"`
}

// ErrorsMessage structure that returns group of errors
type ErrorsMessage struct {
	Status  int         `json:"status"`
	Error   string      `json:"error"`
	Message interface{} `json:"message"`
}

// SuccessfullyMessage structure that returns successfully
type SuccessfullyMessage struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

// Map is a convenient way to create objects of unknown types.
type Map map[string]interface{}

// JSON standardized JSON response.
func JSON(w http.ResponseWriter, r *http.Request, statusCode int, data interface{}) error {
	if data == nil {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(statusCode)
		return nil
	}
	j, err := json.Marshal(data)
	if err != nil {
		return err
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(statusCode)
	_, err = w.Write(j)
	if err != nil{
		return err
	}
	return nil
}

// HTTPError standardized error response in JSON format.
func HTTPError(w http.ResponseWriter, r *http.Request, statusCode int, error string, message string) error {
	msg := ErrorMessage{
		Status:  statusCode,
		Error:   error,
		Message: message,
	}
	return JSON(w, r, statusCode, msg)
}
