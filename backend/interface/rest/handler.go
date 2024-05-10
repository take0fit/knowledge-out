package rest

import (
	"github.com/go-chi/chi/v5"
	"github.com/take0fit/knowledge-out/interface/rest/controller"
	"net/http"
)

type Handler struct {
	Router         *chi.Mux
	AuthController *controller.AuthController
}

func NewHandler(authController *controller.AuthController) *Handler {
	router := chi.NewRouter()

	// API version prefix
	router.Route("/v1/api", func(r chi.Router) {
		r.Get("/login", authController.Login)
		r.Get("/callback", authController.Callback)
	})

	return &Handler{
		Router:         router,
		AuthController: authController,
	}
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.Router.ServeHTTP(w, r)
}
