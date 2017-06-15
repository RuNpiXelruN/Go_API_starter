package utils

import (
	"encoding/json"
	"net/http"
)

// converts db calls to json for client
func JSON(w http.ResponseWriter, val interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	data, err := json.Marshal(val)
	if err != nil {
		panic("Error converting to json :(")
	}

	w.Write(data)
}
