package alert

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func AlertCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		ctx := context.WithValue(r.Context(), "alertID", id)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
