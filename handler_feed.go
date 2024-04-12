package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/manelfitouri/Rssagg/internal/database"
)

// handlerUsersCreate est un gestionnaire HTTP pour la création d'un nouvel utilisateur
func (cfg *apiConfig) handlerCreateFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	// Structure pour les paramètres de la requête
	type parameters struct {
		Name string
		URL  string `json:"url"`
	}

	// Décodeur JSON pour lire les données de la requête
	decoder := json.NewDecoder(r.Body)
	params := parameters{}

	// Décodage des paramètres de la requête
	err := decoder.Decode(&params)
	if err != nil {
		// En cas d'erreur lors du décodage, répondre avec une erreur interne
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters")
		return
	}

	// Création d'un nouvel utilisateur dans la base de données
	feed, err := cfg.DB.Createfeed(r.Context(), database.CreatefeedParams{
		ID:        uuid.New(),       // Génération d'un nouvel identifiant UUID
		CreatedAt: time.Now().UTC(), // Définition de la date et heure de création
		UpdatedAt: time.Now().UTC(), // Définition de la date et heure de mise à jour
		UserID:    user.ID,
		Name:      params.Name,
		Url:       params.URL,
	})
	if err != nil {
		// En cas d'erreur lors de la création du feed, répondre avec une erreur 400
		respondWithError(w, 400, "Couldn't create feed")
		return
	}

	// Répondre avec l'utilisateur créé au format JSON et le code de statut 201 (created)
	respondWithJSON(w, 201, feed)
}
