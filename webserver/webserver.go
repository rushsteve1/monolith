package webserver

import (
	"context"
	"database/sql"
	"embed"
	"fmt"

	log "github.com/sirupsen/logrus"
	"rushsteve1.us/monolith/shared"
)

//go:embed templates/*.html
var templatesFS embed.FS
var loadedTemplates *shared.TemplateHelper

type WebServer struct {
	Config   shared.Config
	Fcgi     bool
	Database *sql.DB
}

func (ws *WebServer) Serve(ctx context.Context) error {
	var err error
	loadedTemplates, err = shared.LoadTemplates(templatesFS, "templates", "layout.html")
	if err != nil {
		return err
	}

	if ws.Database != nil {
		err = createTables(ws.Database, ctx)
		if err != nil {
			return err
		}
	} else {
		log.Fatal("No database given to ", ws.Name())
	}

	return shared.ServeHelper(GetMux(ws, ctx), ws)
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
