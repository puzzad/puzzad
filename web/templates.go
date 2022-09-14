package web

import (
	"errors"
	"html/template"
	"io/fs"
	"net/http"
	"os"
	"path/filepath"

	"github.com/fsnotify/fsnotify"
	"github.com/rs/zerolog/log"
)

func getWatcher() *fsnotify.Watcher {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal().Err(err).Msg("Unable to create template watcher")
	}
	return watcher
}

func (web *Webserver) startWatcher(watcher *fsnotify.Watcher) {
	go func() {
		for {
			select {
			case _, ok := <-watcher.Events:
				if !ok {
					return
				}
				dfs := os.DirFS(filepath.Join("web", "resources", "templates"))
				templates, err := template.ParseFS(dfs, "*.gohtml")
				if err != nil {
					log.Error().Err(err).Msg("Unable to recompile templates")
				} else {
					web.templates = templates
				}
			}
		}
	}()
	err := watcher.Add(filepath.Join("web", "resources", "templates"))
	if err != nil {
		log.Error().Err(err).Msg("Unable to watch templates path")
	}
}

func (web *Webserver) getTemplates() *template.Template {
	_, err := os.OpenFile(filepath.Join("web", "resources", "templates"), os.O_RDONLY, 0644)
	if errors.Is(err, os.ErrNotExist) {
		pfs, _ := fs.Sub(publicfs, "resources/templates")
		return template.Must(template.ParseFS(pfs, "*.gohtml"))
	} else {
		web.startWatcher(getWatcher())
		dfs := os.DirFS(filepath.Join("web", "resources", "templates"))
		return template.Must(template.ParseFS(dfs, "*.gohtml"))
	}
}

func (web *Webserver) handleTemplate(templateName string, data interface{}) func(writer http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		err := web.templates.ExecuteTemplate(writer, templateName+".gohtml", struct {
			Authed bool
			Values map[string][]string
			Data   interface{}
			Admin  bool
		}{
			Authed: web.sessionSore.GetString(request, "username") != "",
			Admin:  web.sessionSore.GetBool(request, "admin"),
			Values: request.URL.Query(),
			Data:   data,
		})
		if err != nil {
			log.Error().Err(err).Msg("Unable to serve index page")
		}
	}
}

func outputError(t *template.Template, writer http.ResponseWriter, code int, message string) {
	err := t.ExecuteTemplate(writer, "error.gohtml", message)
	if err != nil {
		log.Error().Err(err).Msg("Unable to serve error page")
	}
	writer.WriteHeader(code)
}
