package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/manelfitouri/Rssagg/internal/database"
)

// handlerUsersCreate est un gestionnaire HTTP pour la création d'un nouvel utilisateur
func (cfg *apiConfig) handlerUsersCreate(w http.ResponseWriter, r *http.Request) {
	// Structure pour les paramètres de la requête
	type parameters struct {
		Name string // Nom de l'utilisateur
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
	user, err := cfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),       // Génération d'un nouvel identifiant UUID
		CreatedAt: time.Now().UTC(), // Définition de la date et heure de création
		UpdatedAt: time.Now().UTC(), // Définition de la date et heure de mise à jour
		Name:      params.Name,      // Récupération du nom de l'utilisateur depuis les paramètres de la requête
	})
	if err != nil {
		// En cas d'erreur lors de la création de l'utilisateur, répondre avec une erreur 400
		respondWithError(w, 400, "Couldn't create user")
		return
	}

	// Répondre avec l'utilisateur créé au format JSON et le code de statut 201 (created)
	respondWithJSON(w, 201, databaseUserToUser(user))
}

// handlerUsersGet est un gestionnaire HTTP pour récupérer les informations d'un utilisateur
func (cfg *apiConfig) handlerUsersGet(w http.ResponseWriter, r *http.Request, user database.User) {
	// Répondre avec les informations de l'utilisateur au format JSON et le code de statut 200 (OK)
	respondWithJSON(w, 200, databaseUserToUser(user))
}
