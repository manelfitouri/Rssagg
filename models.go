package main

import (
	"time"

	"github.com/google/uuid"                           // Import du package uuid pour générer des identifiants uniques
	"github.com/manelfitouri/Rssagg/internal/database" // Import du package database interne
)

// User est une structure représentant un utilisateur de l'application
type User struct {
	ID        uuid.UUID `json:"id"`         // ID est l'identifiant unique de l'utilisateur
	CreatedAt time.Time `json:"created_at"` // CreatedAt est la date et l'heure de création de l'utilisateur
	UpdatedAt time.Time `json:"updated_at"` // UpdatedAt est la date et l'heure de la dernière mise à jour de l'utilisateur
	Name      string    `json:"name"`       // Name est le nom de l'utilisateur
	ApiKey    string    `json:"api_key"`    // ApiKey est la clé API de l'utilisateur
}

// databaseUserToUser est une fonction qui convertit un utilisateur de la base de données en un utilisateur de l'application
func databaseUserToUser(user database.User) User {
	return User{ // Retourne une instance de User avec les données de l'utilisateur de la base de données
		ID:        user.ID,        // Copie l'ID de l'utilisateur de la base de données
		CreatedAt: user.CreatedAt, // Copie la date de création de l'utilisateur de la base de données
		UpdatedAt: user.UpdatedAt, // Copie la date de mise à jour de l'utilisateur de la base de données
		Name:      user.Name,      // Copie le nom de l'utilisateur de la base de données
		ApiKey:    user.ApiKey,    // Copie la clé API de l'utilisateur de la base de données
	}
}
