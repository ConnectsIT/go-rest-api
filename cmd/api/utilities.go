package main

import (
	"encoding/json"
	"net/http"
)

func (app *application) writeJson(w http.ResponseWriter, status int, data interface{}, wrap string) error {
	wrapper := make(map[string]interface{})

	wrapper[wrap] = data
	json, err := json.Marshal(wrapper)

	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(json)

	return nil

}

func (app *application) errorJson(w http.ResponseWriter, err error) {

	type jsonError struct {
		message string `json:"message"`
	}

	theError := jsonError{
		message: err.Error(),
	}

	app.writeJson(w, http.StatusBadRequest, theError, "error")

}
