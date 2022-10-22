package webserver

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"

	"rushsteve1.us/monolith/shared"
)

type WebServer struct {
	Config   shared.Config
	Fcgi     bool
	Database *sql.DB
}

func (ws *WebServer) Serve(ctx context.Context) error {
	var err error
	err = loadTemplates()
	if err != nil {
		return err
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/",
		func(w http.ResponseWriter, r *http.Request) {
			err := templates.ExecuteTemplate(w, "index.html", nil)
			if err != nil {
				http.Error(w, err.Error(), 500)
			}
		})

	return shared.ServeHelper(mux, ws)
}

func (ws WebServer) Addr() string {
	return ws.Config.WebServer.Addr
}

func (ws WebServer) Name() string {
	return "WebServer"
}

func (ws WebServer) UseFcgi() bool {
	return ws.Fcgi
}

func (ws WebServer) String() string {
	return fmt.Sprintf("%s on %s", ws.Name(), ws.Addr())
}
