-- name: CreateUser :one
-- Requête SQL pour insérer un nouvel utilisateur dans la table "users"
-- La clé API est générée automatiquement en utilisant la fonction sha256 et random() de PostgreSQL
-- La clé générée est encodée en hexadécimal avant d'être stockée dans la base de données
INSERT INTO users (id, created_at, updated_at, name, api_key)
VALUES (
    $1,                     -- Paramètre 1 : ID de l'utilisateur
    $2,                     -- Paramètre 2 : Date de création
    $3,                     -- Paramètre 3 : Date de mise à jour
    $4,                     -- Paramètre 4 : Nom de l'utilisateur
    encode(sha256(random()::text::bytea), 'hex')   -- Générer une clé API aléatoire
)
RETURNING *;               -- Renvoyer toutes les colonnes de l'utilisateur nouvellement créé

-- name: GetUserByAPIKey :one
-- Requête SQL pour récupérer un utilisateur à partir de sa clé API
SELECT * FROM users WHERE api_key = $1;  -- Recherche de l'utilisateur avec la clé API spécifiée
