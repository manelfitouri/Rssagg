package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// respondWithError fonction pour répondre avec une erreur HTTP au format JSON
func respondWithError(w http.ResponseWriter, code int, msg string) {
	// Si le code d'erreur est supérieur à 499, on le log
	if code > 499 {
		log.Printf("Responding with 5XX error: %s", msg)
	}

	// Définition de la structure pour la réponse d'erreur au format JSON
	type errorResponse struct {
		Error string `json:"error"` // Champ pour le message d'erreur
	}

	// Appel de la fonction respondWithJSON pour renvoyer la réponse au client avec le code d'erreur et le message JSON
	respondWithJSON(w, code, errorResponse{
		Error: msg,
	})
}

// respondWithJSON fonction pour répondre avec un corps de réponse au format JSON
func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	// Marshaling du payload en JSON
	data, err := json.Marshal(payload) // Convertit les données en JSON
	if err != nil {
		// Si une erreur se produit lors du marshaling, on log l'erreur et on renvoie une erreur 500 au client
		log.Printf("failed to marshal JSON response: %v", payload)
		w.WriteHeader(500) // Erreur serveur interne
		return
	}

	// Ajout d'un en-tête à la réponse pour s'assurer que la réponse est au format JSON
	w.Header().Add("content-type", "application/json")
	// Définition du code de statut HTTP dans l'en-tête de la réponse
	w.WriteHeader(code)
	// Écriture des données JSON dans le corps de la réponse
	w.Write(data)
}
