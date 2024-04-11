-- +goose Up
-- Requête SQL exécutée lors de la montée de version de la base de données
CREATE TABLE users (
    id UUID PRIMARY KEY,         -- Champ ID de type UUID, clé primaire
    created_at TIMESTAMP NOT NULL,   -- Champ created_at de type TIMESTAMP, non nul
    updated_at TIMESTAMP NOT NULL,   -- Champ updated_at de type TIMESTAMP, non nul
    name TEXT NOT NULL           -- Champ name de type TEXT, non nul
);

-- +goose Down
-- Requête SQL exécutée lors de la descente de version de la base de données
DROP TABLE users;   -- Supprime la table "users"
