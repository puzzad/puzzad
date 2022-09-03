package web

import (
	"context"
	"embed"
	"fmt"
	"io/fs"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/httprate"
	"github.com/rs/zerolog"
)

//go:embed public
var publicfs embed.FS

type Webserver struct {
	handler *http.Server
	router  *chi.Mux
	log     *zerolog.Logger
}

func (web *Webserver) Init(port int, log *zerolog.Logger) {
	web.router = chi.NewRouter()
	web.log = log
	web.addMiddleWare()
	web.addRoutes()
	web.handler = &http.Server{
		Addr:    fmt.Sprintf("%s:%d", "0.0.0.0", port),
		Handler: web.router,
	}
}

func (web *Webserver) addMiddleWare() {
	web.router.Use(middleware.RequestID)
	web.router.Use(middleware.RealIP)
	web.router.Use(middleware.Heartbeat("/ping"))
	web.router.Use(middleware.Recoverer)
	web.router.Use(loggerMiddleware(web.log))
	web.router.Use(middleware.Timeout(60 * time.Second))
	web.router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: false,
		MaxAge:           300,
	}))
	web.router.Use(httprate.LimitByIP(100, 1*time.Minute))
}

func (web *Webserver) addRoutes() {
	pfs, _ := fs.Sub(publicfs, "public")
	web.router.Handle("/*", http.FileServer(http.FS(pfs)))
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
