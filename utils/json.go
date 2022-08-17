package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
	js, err := json.Marshal(data)
	if err != nil {
		fmt.Fprintf(w, "%s", err.Error())
	}
	w.WriteHeader(statusCode)
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func Error(w http.ResponseWriter, statusCode int, err error) {
	type jsonError struct {
		Message string `json:"message"`
	}
	theError := jsonError{
		Message: err.Error(),
	}
	JSON(w, statusCode, theError)
}
