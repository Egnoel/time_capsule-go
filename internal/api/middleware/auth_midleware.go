package middleware

import (
	"context"
	"net/http"

	"github.com/alexedwards/scs/v2"
)

type contextKey string

const UserIDKey contextKey = "user_id"

func RequireAuth(sessionManager *scs.SessionManager, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userID := sessionManager.GetString(r.Context(), "user_id")

		if userID == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), UserIDKey, userID)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func GetUserID(ctx context.Context) string {
	userID, _ := ctx.Value(UserIDKey).(string)
	return userID
}
