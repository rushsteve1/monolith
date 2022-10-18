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
	ctx, err := context.WithCancel(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	serv.Serve(ctx)
}

func ServeHelper(mux *http.ServeMux, serv Service) error {
	listener, err := net.Listen("tcp", serv.Addr())
	if err != nil {
		log.Fatalf("Could not start %s listener: %v", serv.Name(), err)
	}

	log.Infof("%s started on %s", serv.Name(), serv.Addr())
	if serv.UseFcgi() {
		fcgi.Serve(listener, mux)
	} else {
		http.Serve(listener, mux)
	}

	return fmt.Errorf("%s exited unexpectedly", serv.Name())
}
