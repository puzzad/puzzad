package web

import (
	"net/http"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/zerolog"
)

func (web *Webserver) AdminOnly(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if web.sessionSore.GetBool(r, "admin") {
			next.ServeHTTP(w, r)
			return
		}
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
	})
}

func (web *Webserver) AccountOnly(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if web.sessionSore.GetString(r, "username") != "" {
			next.ServeHTTP(w, r)
			return
		}
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
	})
}

func (web *Webserver) loggerMiddleware(logger *zerolog.Logger) func(next http.Handler) http.Handler {
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

type Interceptor struct {
	handler    func(string, interface{}) func(writer http.ResponseWriter, request *http.Request)
	writer     http.ResponseWriter
	request    *http.Request
	overridden bool
}

func (i *Interceptor) WriteHeader(rc int) {
	switch rc {
	case 500:
		i.writer.Header().Set("Content-Type", "text/html")
		i.handler("500", nil)(i.writer, i.request)
		return
	case 404:
		i.writer.Header().Set("Content-Type", "text/html")
		i.handler("404", nil)(i.writer, i.request)
		return
	case 403:
		i.writer.Header().Set("Content-Type", "text/html")
		i.handler("403", nil)(i.writer, i.request)
		return
	case 401:
		i.writer.Header().Set("Content-Type", "text/html")
		i.handler("401", nil)(i.writer, i.request)
		return
	default:
		i.writer.WriteHeader(rc)
		return
	}
}

func (i *Interceptor) Write(b []byte) (int, error) {
	if !i.overridden {
		return i.writer.Write(b)
	}
	return 0, nil
}

func (i *Interceptor) Header() http.Header {
	return i.writer.Header()
}

func (web *Webserver) errorInterceptMiddleWare(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		wrapper := &Interceptor{
			handler: web.handleTemplate,
			writer:  w,
			request: r,
		}
		handler.ServeHTTP(wrapper, r)
	})
}
