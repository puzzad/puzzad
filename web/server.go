package web

import (
	"context"
	"embed"
	"errors"
	"fmt"
	"html/template"
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
	"github.com/greboid/puzzad/ent"
	"github.com/greboid/puzzad/puzzad"
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
	sessionSore *sessions.Session
	static      http.Handler
	templates   *template.Template
	watcher     *fsnotify.Watcher

	Client           *database.DBClient
	UserManager      *puzzad.UserManager
	AdventureManager *puzzad.AdventureManager
	SessionKey       string
	EnableLogs       bool
}

func (web *Webserver) Init(port int, log *zerolog.Logger) {
	web.sessionSore = sessions.New([]byte(web.SessionKey))
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
	if web.EnableLogs {
		web.router.Use(web.loggerMiddleware(web.log))
	}
	web.router.Use(middleware.Timeout(60 * time.Second))
	web.router.Use(web.sessionSore.Enable)
	web.router.Use(web.errorInterceptMiddleWare)
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
	web.router.Get("/", web.handleTemplate("index", nil))
	web.router.Get("/register", web.handleTemplate("register", nil))
	web.router.Post("/register", web.handleRegister)
	web.router.Get("/login", web.handleTemplate("login", nil))
	web.router.Post("/login", web.handleLogin)
	// TODO: Handle this better?
	web.router.Get("/logout", web.handleLogout)
	web.router.Post("/logout", web.handleLogout)
	web.router.Get("/passreset", web.handleTemplate("passreset", nil))
	web.router.Post("/passreset", web.handlePassReset)
	web.router.Get("/setpass", web.handleTemplate("setpass", nil))
	web.router.Post("/setpass", web.handleSetPass)
	web.router.With(web.AdminOnly).Mount("/admin", web.adminRoutes())
	web.router.With(web.AccountOnly).Mount("/account", web.accountRoutes())
	web.router.Get("/validate", web.handleTemplate("validate", nil))
	web.router.Post("/validate", web.handleValidate)
	web.router.Get("/adventures", web.handleListAdventures)
	web.router.Get("/play", web.handleTemplate("play", nil))
	web.router.Post("/play", func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(http.StatusInternalServerError)
	})
	web.router.Handle("/*", web.static)
}

func (web *Webserver) adminRoutes() http.Handler {
	r := chi.NewRouter()
	r.Get("/", web.handleTemplate("admin", nil))
	return r
}

func (web *Webserver) accountRoutes() http.Handler {
	r := chi.NewRouter()
	r.Get("/", web.handleTemplate("account", nil))
	return r
}

func (web *Webserver) handleValidate(writer http.ResponseWriter, request *http.Request) {
	code := request.FormValue("code")

	u, err := web.UserManager.CompleteVerification(request.Context(), code)
	if err != nil {
		outputError(web.templates, writer, http.StatusBadRequest, "Incorrect verify code")
		return
	}
	log.Debug().Str("username", u.Email).Msg("Adding Auth: ")
	web.sessionSore.Put(request, "username", u.Email)

	writer.WriteHeader(http.StatusTemporaryRedirect)
	if request.Header.Get("HX-Request") != "" {
		writer.Header().Set("HX-Redirect", "/index.html")
		writer.WriteHeader(http.StatusTemporaryRedirect)
	} else {
		http.Redirect(writer, request, "/index.html", http.StatusTemporaryRedirect)
	}
}

func (web *Webserver) handleLogin(writer http.ResponseWriter, request *http.Request) {
	username := request.FormValue("username")
	password := request.FormValue("password")

	u, err := web.UserManager.Authenticate(request.Context(), username, password)
	if err != nil {
		if errors.Is(err, puzzad.ErrBadUsernameOrPassword) {
			outputError(web.templates, writer, http.StatusUnauthorized, "Invalid username or password")
		} else if errors.Is(err, puzzad.ErrAccountDisabled) {
			outputError(web.templates, writer, http.StatusUnauthorized, "Account disabled")
		} else if errors.Is(err, puzzad.ErrAccountUnverified) {
			outputError(web.templates, writer, http.StatusUnauthorized, "Account unverified")
		} else {
			outputError(web.templates, writer, http.StatusInternalServerError, "Internal server error")
		}
		return
	}

	log.Debug().Str("username", username).Msg("Adding Auth: ")
	web.sessionSore.Put(request, "username", u.Email)
	web.sessionSore.Put(request, "admin", u.Admin)
	if request.Header.Get("HX-Request") != "" {
		writer.Header().Set("HX-Redirect", "index.html")
		writer.WriteHeader(http.StatusTemporaryRedirect)
	} else {
		http.Redirect(writer, request, "index.html", http.StatusTemporaryRedirect)
	}
}

func (web *Webserver) handleRegister(writer http.ResponseWriter, request *http.Request) {
	email := request.FormValue("email")
	password := request.FormValue("password")
	u, err := web.UserManager.CreateUser(request.Context(), email)
	if err != nil {
		outputError(web.templates, writer, http.StatusInternalServerError, "Internal server error")
		return
	}
	err = web.UserManager.SetPassword(request.Context(), u, password)
	if err != nil {
		outputError(web.templates, writer, http.StatusInternalServerError, "Internal server error")
		return
	}
	if request.Header.Get("HX-Request") != "" {
		writer.Header().Set("HX-Redirect", "/validate")
		writer.WriteHeader(http.StatusTemporaryRedirect)
	} else {
		http.Redirect(writer, request, "/validate", http.StatusTemporaryRedirect)
	}
}

func (web *Webserver) handlePassReset(writer http.ResponseWriter, request *http.Request) {
	email := request.FormValue("email")
	_ = web.UserManager.StartPasswordReset(request.Context(), email)
	if request.Header.Get("HX-Request") != "" {
		_, _ = writer.Write([]byte("<p>If an account exists for that email address, an email will arrive shortly with further instructions.</p>"))
	} else {
		http.Redirect(writer, request, "/", http.StatusTemporaryRedirect)
	}
}

func (web *Webserver) handleSetPass(writer http.ResponseWriter, request *http.Request) {
	code := request.FormValue("code")
	password := request.FormValue("password")
	confirm := request.FormValue("confirm")
	if password != confirm {
		outputError(web.templates, writer, http.StatusBadRequest, "passwords do not match")
		return
	}
	success, _ := web.UserManager.FinishPasswordReset(request.Context(), code, password)
	if !success {
		outputError(web.templates, writer, http.StatusInternalServerError, "error resetting password")
		return
	}
	if request.Header.Get("HX-Request") != "" {
		writer.Header().Set("HX-Redirect", "/login")
		writer.WriteHeader(http.StatusTemporaryRedirect)
	} else {
		http.Redirect(writer, request, "/login", http.StatusTemporaryRedirect)
	}
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

func (web *Webserver) handleListAdventures(writer http.ResponseWriter, request *http.Request) {
	as, err := web.AdventureManager.GetAllPublicAdventures(request.Context())
	if err != nil {
		// TODO Should probably handle this better
		log.Error().Err(err).Msg("Unable to load adventure list")
		as = make([]*ent.Adventure, 0)
	}
	web.handleTemplate("adventures", as)(writer, request)
}
