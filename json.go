package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondWithError(w http.ResponseWriter, code int, msg string) {
	if code > 499 {
		log.Printf("Responding with 5XX error: %s", msg)
	}
	type errorResponse struct {
		Error string `json:"error"`
	}
	respondWithJSON(w, code, errorResponse{
		Error: msg,
	})
}

// all response bodies will have a json format
func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	// rassembler le payload into a json object
	data, err := json.Marshal(payload) //marshal the given data into a json string and return it as bytes
	if err != nil {
		log.Printf("failed to marshal JSON response: %v", payload)
		w.WriteHeader(500) //internal server error
		return
	}
	//add a header to the response to make sure that the response is a json
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(code)
	w.Write(data)

}
