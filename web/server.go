package web

import (
	"context"
	"crypto/rand"
	"embed"
	"encoding/json"
	"errors"
	"fmt"
	"html/template"
	"io"
	"io/fs"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/httprate"
	"github.com/golangcollege/sessions"
	"github.com/greboid/puzzad/puzzad/database"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

//go:embed resources
var publicfs embed.FS

type Webserver struct {
	handler     *http.Server
	router      *chi.Mux
	log         *zerolog.Logger
	Client      *database.DBClient
	sessionSore *sessions.Session
	static      http.Handler
	templates   *template.Template
	watcher     *fsnotify.Watcher
}

type Login struct {
	Code     string
	Username string
	Password string
}

func (web *Webserver) Init(port int, log *zerolog.Logger) {
	// TODO: Pull this from an env var
	web.sessionSore = sessions.New(randomByte(32))
	web.router = chi.NewRouter()
	web.log = log
	web.static = web.getStaticServer()
	web.templates = web.getTemplates()
	web.addMiddleWare()
	web.addRoutes()
	web.handler = &http.Server{
		Addr:    fmt.Sprintf("%s:%d", "0.0.0.0", port),
		Handler: web.router,
	}
}

func getWatcher() *fsnotify.Watcher {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Panic().Err(err).Msg("Unable to create template watcher")
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

func (web *Webserver) getStaticServer() http.Handler {
	_, err := os.OpenFile(filepath.Join("web", "resources", "public"), os.O_RDONLY, 0644)
	if errors.Is(err, os.ErrNotExist) {
		pfs, _ := fs.Sub(publicfs, "resources/public")
		return http.FileServer(http.FS(pfs))
	} else {
		return http.FileServer(http.FS(os.DirFS(filepath.Join("web", "resources", "public"))))
	}
}

func (web *Webserver) addMiddleWare() {
	web.router.Use(middleware.RequestID)
	web.router.Use(middleware.RealIP)
	web.router.Use(middleware.Heartbeat("/ping"))
	web.router.Use(middleware.Recoverer)
	web.router.Use(loggerMiddleware(web.log))
	web.router.Use(middleware.Timeout(60 * time.Second))
	web.router.Use(web.sessionSore.Enable)
	web.router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token", "hx-current-url", "hx-request", "hx-target", "hx-trigger", "hx-trigger-name"},
		AllowCredentials: false,
		MaxAge:           300,
	}))
	web.router.Use(httprate.LimitByIP(100, 1*time.Minute))
}

func (web *Webserver) addRoutes() {
	web.router.Get("/", web.handleTemplate("index"))
	web.router.Get("/register", web.handleTemplate("register"))
	web.router.Post("/register", web.handleRegister)
	web.router.Get("/login", web.handleTemplate("login"))
	web.router.Post("/login", web.handleLogin)
	web.router.Post("/logout", web.handleLogout)
	web.router.Get("/passreset", web.handleTemplate("passreset"))
	web.router.Post("/passreset", func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(http.StatusInternalServerError)
	})
	web.router.Get("/validate", web.handleTemplate("validate"))
	web.router.Post("/validate", web.handleValidate)
	web.router.Get("/play", web.handleTemplate("play"))
	web.router.Post("/play", func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(http.StatusInternalServerError)
	})
	web.router.Handle("/*", web.static)
}

func (web *Webserver) handleValidate(writer http.ResponseWriter, request *http.Request) {
	time.Sleep(2 * time.Second)
	writer.WriteHeader(http.StatusTemporaryRedirect)
	if request.Header.Get("HX-Request") != "" {
		writer.Header().Set("HX-Redirect", "/index.html")
		writer.WriteHeader(http.StatusTemporaryRedirect)
	} else {
		http.Redirect(writer, request, "/index.html", http.StatusTemporaryRedirect)
	}
}

func (web *Webserver) handleLogin(writer http.ResponseWriter, request *http.Request) {
	data := &Login{}
	bytes, err := io.ReadAll(request.Body)
	if err != nil {
		return
	}
	err = json.Unmarshal(bytes, data)
	if err != nil {
		return
	}
	log.Debug().Str("username", data.Username).Msg("Adding Auth: ")
	web.sessionSore.Put(request, "username", data.Username)
	if request.Header.Get("HX-Request") != "" {
		writer.Header().Set("HX-Redirect", "index.html")
		writer.WriteHeader(http.StatusTemporaryRedirect)
	} else {
		http.Redirect(writer, request, "index.html", http.StatusTemporaryRedirect)
	}
}

func (web *Webserver) handleRegister(writer http.ResponseWriter, request *http.Request) {
	time.Sleep(2 * time.Second)
	if request.Header.Get("HX-Request") != "" {
		writer.Header().Set("HX-Redirect", "/validate.html")
		writer.WriteHeader(http.StatusTemporaryRedirect)
	} else {
		http.Redirect(writer, request, "/validate.html", http.StatusTemporaryRedirect)
	}
}

func (web *Webserver) RunAndWait() error {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM)
	signal.Notify(stop, syscall.SIGINT)
	var servErr error
	go func() {
		if err := web.handler.ListenAndServe(); err != nil {
			if err != http.ErrServerClosed {
				servErr = err
			}
			close(stop)
		}
	}()
	<-stop
	if servErr != nil {
		return servErr
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return web.handler.Shutdown(ctx)
}

func (web *Webserver) handleLogout(writer http.ResponseWriter, request *http.Request) {
	log.Debug().Str("username", web.sessionSore.GetString(request, "username")).Msg("Removing Auth: ")
	web.sessionSore.Destroy(request)
	if request.Header.Get("HX-Request") != "" {
		writer.Header().Set("HX-Redirect", "index.html")
		writer.WriteHeader(http.StatusTemporaryRedirect)
	} else {
		http.Redirect(writer, request, "index.html", http.StatusTemporaryRedirect)
	}
}

func (web *Webserver) handleTemplate(templateName string) func(writer http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		err := web.templates.ExecuteTemplate(writer, templateName+".gohtml", struct {
			Authed bool
		}{
			Authed: web.sessionSore.GetString(request, "username") != "",
		})
		if err != nil {
			log.Error().Err(err).Msg("Unable to serve index page")
		}
	}
}

func loggerMiddleware(logger *zerolog.Logger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			wrapper := middleware.NewWrapResponseWriter(w, r.ProtoMajor)
			defer func() {
				logger.Info().
					Timestamp().
					Str("REMOTE_ADDR", r.RemoteAddr).
					Str("url", r.URL.Path).
					Str("Proto", r.Proto).
					Str("Method", r.Method).
					Str("User-Agent", r.Header.Get("User-Agent")).
					Int("status", wrapper.Status()).
					Msg("incoming_request")
			}()
			next.ServeHTTP(wrapper, r)
		})
	}
}

func randomByte(length int) []byte {
	key := make([]byte, length)

	_, err := rand.Read(key)
	if err != nil {
		return nil
	}
	return key
}
