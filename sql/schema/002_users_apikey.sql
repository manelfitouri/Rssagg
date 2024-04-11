-- +goose Up
-- Requête SQL exécutée lors de la montée de version de la base de données
-- Ajoute une nouvelle colonne "api_key" à la table "users"
-- La colonne est de type VARCHAR(64), unique et non nulle
-- La valeur par défaut est générée en utilisant la fonction sha256 et random() de PostgreSQL
ALTER TABLE users ADD COLUMN api_key VARCHAR(64) UNIQUE NOT NULL DEFAULT (
    encode(sha256(random()::text::bytea), 'hex')
);

-- +goose Down
-- Requête SQL exécutée lors de la descente de version de la base de données
-- Supprime la colonne "api_key" de la table "users"
ALTER TABLE users DROP COLUMN api_key;
