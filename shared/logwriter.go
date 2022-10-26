package shared

import (
	"net/http"

	log "github.com/sirupsen/logrus"
)

type LogWriter struct {
	name string
	r    *http.Request
	w    http.ResponseWriter
}

func (lw LogWriter) Header() http.Header {
	return lw.w.Header()
}

func (lw LogWriter) Write(b []byte) (int, error) {
	return lw.w.Write(b)
}
func (lw LogWriter) WriteHeader(statusCode int) {
	log.WithFields(
		log.Fields{
			"url":     lw.r.URL,
			"method":  lw.r.Method,
			"remote":  lw.r.RemoteAddr,
			"service": lw.name,
			"status":  statusCode,
		}).Info("HTTP Request")
	lw.w.WriteHeader(statusCode)
}

func LogWrapper(handler http.Handler, serv Service) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		lw := LogWriter{name: serv.Name(), w: w, r: r}
		handler.ServeHTTP(lw, r)
	})
}
