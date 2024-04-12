package main

import (
	"net/http"

	"github.com/manelfitouri/Rssagg/internal/database"
	"github.com/manelfitouri/Rssagg/internal/database/auth"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (cfg *apiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetAPIKey(r.Header)
		if err != nil {
			// Si la clé API n'est pas trouvée, répondre avec une erreur 403 (forbidden)
			respondWithError(w, 403, "Couldn't find api key")
			return
		}

		// Récupération des informations de l'utilisateur à partir de la base de données en utilisant la clé API
		user, err := cfg.DB.GetUserByAPIKey(r.Context(), apiKey)
		if err != nil {
			// En cas d'erreur lors de la récupération des informations de l'utilisateur, répondre avec une erreur 400
			respondWithError(w, 400, "Couldn't get user")
			return
		}

		handler(w, r, user)
	}
}
