module github.com/manelfitouri/Rssagg // Définition du module avec son chemin d'importation

go 1.22.2 // Version minimale de Go requise pour le projet

// Dépendances requises par le projet
require github.com/joho/godotenv v1.5.1 // Dépendance pour la gestion des variables d'environnement depuis un fichier .env (indirecte)
require github.com/go-chi/chi/v5 v5.0.12 // Dépendance pour le routeur HTTP Chi version 5
require (
    github.com/go-chi/cors v1.2.1 // Dépendance pour le middleware CORS pour Chi
    github.com/google/uuid v1.6.0 // Dépendance pour la génération d'identifiants UUID
)
require github.com/lib/pq v1.10.9 // Dépendance pour le pilote PostgreSQL
