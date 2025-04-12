package events

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
)

// lire {id} depuis l'URL et le met dans le context
func EventCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id") // récupère la valeur {id}
		ctx := context.WithValue(r.Context(), "eventID", id)
		next.ServeHTTP(w, r.WithContext(ctx)) // continue la requête avec ce context
	})
}
