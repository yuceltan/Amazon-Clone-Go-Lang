package routes

import (
	"log"
	"net/http"

	"bitbucket.org/janpavtel/site/internal/models"
	"bitbucket.org/janpavtel/site/internal/tokens"
)

func AuthenticationValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		enrichedRequest, authenticated := tokens.IsAuthenticated(r)
		if ! authenticated {
			log.Println("User not authenticated")
			http.Redirect(w, r, "/users/login", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, enrichedRequest)
	})
}

func AdminRoleValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		enrichedRequest, authenticated := tokens.IsAuthenticated(r, models.ROLE_ADMIN)
		if ! authenticated {
			log.Println("User don't have admin role")
			http.Redirect(w, r, "/users/login", http.StatusForbidden)
			return
		}

		next.ServeHTTP(w, enrichedRequest)
	})
}

func AdminOrEditorRoleValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		enrichedRequest, authenticated := tokens.IsAuthenticated(r, models.ROLE_ADMIN, models.ROLE_EDITOR)
		if ! authenticated {
			log.Println("User don't have admin or editor role")
			http.Redirect(w, r, "/users/login", http.StatusForbidden)
			return
		}

		next.ServeHTTP(w, enrichedRequest)
	})
}

func UpdateJWT(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		err := tokens.UpdateJWTandSaveToCookie(w, r)
		if err != nil {
			log.Println("Failed to update token ", err)
			http.Redirect(w, r, "/users/login", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}