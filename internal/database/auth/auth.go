package auth

import (
	"errors"
	"net/http"
	"strings"
)

// GetAPIKey extrait une clé API des en-têtes d'une requête HTTP
// ex: Authorization: ApiKey {clé API ici}
func GetAPIKey(headers http.Header) (string, error) {
	// Récupération de la valeur de l'en-tête "Authorization"
	val := headers.Get("Authorization")
	if val == "" {
		// Si aucun en-tête "Authorization" n'est trouvé, retourner une erreur
		return "", errors.New("no authentication info found")
	}
	// Séparation de la valeur de l'en-tête en deux parties sur l'espace
	vals := strings.Split(val, " ")
	if len(vals) != 2 {
		// Si la valeur de l'en-tête ne contient pas deux parties, retourner une erreur
		return "", errors.New("malformed authorization header")
	}
	if vals[0] != "ApiKey" {
		// Si la première partie de la valeur de l'en-tête n'est pas "ApiKey", retourner une erreur
		return "", errors.New("malformed authorization header")
	}
	// Retourner la deuxième partie de la valeur de l'en-tête (la clé API) et aucune erreur
	return vals[1], nil
}
