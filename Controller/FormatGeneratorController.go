package controller

import (
	"encoding/json"
	"net/http"
)

// var VOWELS String = "AEIOU"

func Health() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		json.NewEncoder(w).Encode("UP and Tired")
	}
}
