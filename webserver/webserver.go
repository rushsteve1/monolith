package webserver

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"rushsteve1.us/monolith/shared"
)

const Name = "WebServer"

type WebServer struct {
	Config shared.Config
	Fcgi   bool
}

func (ws *WebServer) Serve(ctx context.Context) error {
	mux := http.NewServeMux()
	mux.HandleFunc("/",
		func(w http.ResponseWriter, r *http.Request) {
			io.WriteString(w, "Hello World\n")
		})

	return shared.ServeHelper(mux, ws)
}

func (ws *WebServer) Addr() string {
	return ws.Config.WebServer.Addr
}

func (ws *WebServer) Name() string {
	return Name
}

func (ws *WebServer) UseFcgi() bool {
	return ws.Fcgi
}

func (ws *WebServer) String() string {
	return fmt.Sprintf("%s on %s", ws.Name(), ws.Addr())
}
