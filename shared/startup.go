package shared

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"net/http/fcgi"

	log "github.com/sirupsen/logrus"
	"github.com/thejerf/suture/v4"
)

func MainHelper(serv suture.Service, name string) {
	ctx, cancel := context.WithCancel(context.Background())
	err := serv.Serve(ctx)

	log.Errorf("%s exited unexpectedly: %v", name, err)
	cancel()
}

func ServeHelper(mux *http.ServeMux, serv Service) error {
	listener, err := net.Listen("tcp", serv.Addr())
	if err != nil {
		log.Fatalf("Could not start %s listener: %v", serv.Name(), err)
	}

	handler := LogWrapper(mux, serv)

	log.Infof("%s started on %s", serv.Name(), serv.Addr())
	if serv.UseFcgi() {
		fcgi.Serve(listener, handler)
	} else {
		http.Serve(listener, handler)
	}

	return fmt.Errorf("%s exited unexpectedly", serv.Name())
}

func LogWrapper(handler http.Handler, serv Service) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.WithFields(log.Fields{"url": r.URL, "method": r.Method, "remote": r.RemoteAddr, "service": serv.Name()}).Info("HTTP Request")
		handler.ServeHTTP(w, r)
	})
}
