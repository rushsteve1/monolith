package webserver

import (
	"context"
	"database/sql"
	"fmt"
	stdlog "log"
	"net/http"
	"net/http/cgi"
	"path/filepath"

	log "github.com/sirupsen/logrus"
	"rushsteve1.us/monolith/shared"
)

type WebServer struct {
	Config   shared.Config
	Fcgi     bool
	Database *sql.DB
}

func (ws *WebServer) Serve(ctx context.Context) error {
	err := loadTemplates()
	if err != nil {
		return err
	}

	err = createTables(ws.Database, ctx)
	if err != nil {
		return err
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/",
		func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path != "/" {
				http.NotFound(w, r)
				return
			}

			err := templates.ExecuteTemplate(w, "index.html", nil)
			if err != nil {
				http.Error(w, err.Error(), 500)
			}
		})

	mux.Handle(
		"/cgi-bin/",
		http.StripPrefix("/cgi-bin/",
			http.HandlerFunc(
				func(w http.ResponseWriter, r *http.Request) {
					path := filepath.Join(ws.Config.WebServer.CgiPath, r.URL.Path)
					l := stdlog.New(log.StandardLogger().WriterLevel(log.ErrorLevel), "", 0)
					h := &cgi.Handler{Path: path, Root: "/cgi-bin/", Logger: l}
					log.Trace("Running CGI script ", path)
					h.ServeHTTP(w, r)
				})))

	if !ws.Config.UseCaddy {
		log.Info("Serving static files without Caddy")
		mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(ws.Config.WebServer.StaticPath))))
	}

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
