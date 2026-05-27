package api

import (
	"net/http"

	"github.com/alexedwards/scs/v2"
	"github.com/go-chi/chi/v5"

	"github.com/egnoel/future-message-go/internal/api/handlers"
	"github.com/egnoel/future-message-go/internal/api/middleware"
)

func NewRouter(
	authHandler *handlers.AuthHandler,
	//	letterHandler *handlers.LetterHandler,
	sessionManager *scs.SessionManager,
) http.Handler {

	r := chi.NewRouter()

	// rotas públicas
	r.Post("/auth/register", authHandler.Register)
	r.Post("/auth/login", authHandler.Login)
	r.Post("/auth/logout", authHandler.Logout)

	// rotas protegidas
	r.Group(func(protected chi.Router) {

		protected.Use(func(next http.Handler) http.Handler {
			return middleware.RequireAuth(sessionManager, next)
		})

		//	protected.Post("/letters", letterHandler.Create)
		//	protected.Get("/letters", letterHandler.List)
	})

	return sessionManager.LoadAndSave(r)
}
