package web

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/greboid/puzzad/ent"
	"github.com/rs/zerolog/log"
)

func (web *Webserver) adminRoutes() http.Handler {
	r := chi.NewRouter()
	r.Get("/", web.handleTemplate("admin", nil))
	r.Get("/adventures", web.handleAdminAdventureList)
	r.Get("/adventures/{id:[0-9]+}", web.handleAdminAdventureDetails)
	r.Post("/adventures", web.handleAdminAdventureCreate)
	r.Patch("/adventures", web.handleAdminAdventureUpdate)
	r.Post("/puzzles", web.handleAdminPuzzlesCreate)
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
		log.Debug().Err(err).Msg("Incorrect adventure ID")
		http.Error(writer, "Unable to find adventure", http.StatusNotFound)
		return
	}
	adv, p, err := web.AdventureManager.GetAdventureByID(request.Context(), id)
	if err != nil {
		if ent.IsNotFound(err) {
			log.Debug().Err(err).Int("id", id).Msg("Unable to find adventure")
			http.Error(writer, "Unable to find adventure", http.StatusNotFound)
		} else {
			log.Debug().Err(err).Msg("Unable to get adventure")
			http.Error(writer, "Unable to get adventure", http.StatusInternalServerError)
		}
		return
	}
	web.handleTemplate("admin-adventure", struct {
		Adventure *ent.Adventure
		Puzzles   []*ent.Puzzle
	}{
		Adventure: adv,
		Puzzles:   p,
	})(writer, request)
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

func (web *Webserver) handleAdminPuzzlesCreate(writer http.ResponseWriter, request *http.Request) {
	aid, err := strconv.Atoi(request.FormValue("aid"))
	name := request.FormValue("puzzleName")
	if err != nil {
		log.Debug().Err(err).Msg("Bad adventure ID")
		outputError(web.templates, writer, http.StatusBadRequest, "Bad adventure ID")
		return
	}
	_, err = web.AdventureManager.CreatePuzzle(request.Context(), aid, name)
	if err != nil {
		log.Debug().Err(err).Msg("Unable to create puzzle")
		outputError(web.templates, writer, http.StatusInternalServerError, "Unable to create puzzle")
		return
	}
	if request.Header.Get("HX-Request") != "" {
		writer.Header().Set("HX-Redirect", fmt.Sprintf("%s%d", "/admin/adventures/", aid))
		writer.WriteHeader(http.StatusTemporaryRedirect)
	} else {
		http.Redirect(writer, request, fmt.Sprintf("%s%d", "/admin/adventures/", aid), http.StatusTemporaryRedirect)
	}
}

func (web *Webserver) handleAdminAdventureUpdate(writer http.ResponseWriter, request *http.Request) {
	aid := request.FormValue("aid")
	name := request.FormValue("name")
	desc := request.FormValue("desc")
	price := request.FormValue("price")
	puzzles := request.FormValue("puzzles")
	// TODO: Actually update this
	log.Debug().Str("aid", aid).Str("name", name).Str("desc", desc).Str("price", price).Str("puzzles", puzzles).Msg("Values:")
}
