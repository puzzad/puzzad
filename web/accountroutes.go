package web

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (web *Webserver) accountRoutes() http.Handler {
	r := chi.NewRouter()
	r.Get("/", web.handleTemplate("account", nil))
	return r
}
