package web

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func (web *Webserver) adminRoutes() http.Handler {
	r := chi.NewRouter()
	r.Get("/", web.handleTemplate("admin", nil))
	r.Get("/adventures", web.handleAdminAdventureList)
	r.Get("/adventures/{id:[0-9]+}", web.handleAdminAdventureDetails)
	r.Post("/adventures", web.handleAdminAdventureCreate)
	return r
}

func (web *Webserver) handleAdminAdventureList(writer http.ResponseWriter, request *http.Request) {
	advs, err := web.AdventureManager.GetAllAdventures(request.Context())
	if err != nil {
		http.Error(writer, "Unable to list adventures", http.StatusInternalServerError)
	}
	web.handleTemplate("admin-adventures", advs)(writer, request)
}

func (web *Webserver) handleAdminAdventureDetails(writer http.ResponseWriter, request *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(request, "id"))
	if err != nil {
		http.Error(writer, "Unable to find adventure", http.StatusNotFound)
	}
	adv, p, err := web.AdventureManager.GetAdventureByID(request.Context(), id)
	if err != nil {
		http.Error(writer, "Unable to get adventure", http.StatusInternalServerError)
	}
	web.handleTemplate("admin-adventure", map[string]interface{}{"Adventure": adv, "Puzzles": p})(writer, request)
}

func (web *Webserver) handleAdminAdventureCreate(writer http.ResponseWriter, request *http.Request) {
	name := request.FormValue("name")
	a, err := web.AdventureManager.CreateAdventure(request.Context(), name)
	if err != nil {
		outputError(web.templates, writer, http.StatusInternalServerError, "Unable to create adventure")
		return
	}
	if request.Header.Get("HX-Request") != "" {
		writer.Header().Set("HX-Redirect", fmt.Sprintf("%s%d", "/admin/adventures/", a.ID))
		writer.WriteHeader(http.StatusTemporaryRedirect)
	} else {
		http.Redirect(writer, request, "/login", http.StatusTemporaryRedirect)
	}
}
