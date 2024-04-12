-- name: Createfeed :one
-- Requête SQL pour insérer un nouveau feed dans la table "feeds"
INSERT INTO feeds (id, created_at, updated_at, name, url, user_id)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;               -- Renvoyer toutes les colonnes de feed nouvellement créé