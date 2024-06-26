// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: users.sql

package database

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (id, created_at, updated_at, name, api_key)
VALUES (
    $1,                     -- Paramètre 1 : ID de l'utilisateur
    $2,                     -- Paramètre 2 : Date de création
    $3,                     -- Paramètre 3 : Date de mise à jour
    $4,                     -- Paramètre 4 : Nom de l'utilisateur
    encode(sha256(random()::text::bytea), 'hex')   -- Générer une clé API aléatoire
)
RETURNING id, created_at, updated_at, name, api_key
`

type CreateUserParams struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string
}

// Requête SQL pour insérer un nouvel utilisateur dans la table "users"
// La clé API est générée automatiquement en utilisant la fonction sha256 et random() de PostgreSQL
// La clé générée est encodée en hexadécimal avant d'être stockée dans la base de données
func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser,
		arg.ID,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.Name,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Name,
		&i.ApiKey,
	)
	return i, err
}

const getUserByAPIKey = `-- name: GetUserByAPIKey :one

SELECT id, created_at, updated_at, name, api_key FROM users WHERE api_key = $1
`

// Renvoyer toutes les colonnes de l'utilisateur nouvellement créé
// Requête SQL pour récupérer un utilisateur à partir de sa clé API
func (q *Queries) GetUserByAPIKey(ctx context.Context, apiKey string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByAPIKey, apiKey)
	var i User
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Name,
		&i.ApiKey,
	)
	return i, err
}
