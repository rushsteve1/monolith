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
	config shared.Config
	dbConn *sql.Conn
}

func New(ctx context.Context, cfg shared.Config, db *sql.DB) *WebServer {
	conn, err := db.Conn(ctx)
	if err != nil {
		log.Fatal(err)
	}
	return &WebServer{config: cfg, dbConn: conn}
}

func (ws *WebServer) Serve(ctx context.Context) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	var err error
	loadedTemplates, err = shared.LoadTemplates(templatesFS, "templates", "layout.html")
	if err != nil {
		return err
	}

	if ws.dbConn != nil {
		err = createTables(ws.dbConn, ctx)
		if err != nil {
			return err
		}
	} else {
		log.Fatal("No database given to ", ws.Name())
	}

	defer ws.dbConn.Close()
	return shared.ServeHelper(GetMux(ws, ctx), ws)
}

func (ws WebServer) Addr() string {
	return ws.config.WebServer.Addr
}

func (ws WebServer) Name() string {
	return "WebServer"
}

func (ws WebServer) UseFcgi() bool {
	return ws.config.UseFcgi
}

func (ws WebServer) String() string {
	return fmt.Sprintf("%s on %s", ws.Name(), ws.Addr())
}

func (ws WebServer) DBConn() *sql.Conn {
	return ws.dbConn
}
