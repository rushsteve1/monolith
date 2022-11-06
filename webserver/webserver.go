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
	db     *sql.DB
}

func New(ctx context.Context, cfg shared.Config, db *sql.DB) *WebServer {
	return &WebServer{config: cfg, db: db}
}

func (ws *WebServer) Serve(ctx context.Context) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	var err error
	loadedTemplates, err = shared.LoadTemplates(templatesFS, "templates", "layout.html")
	if err != nil {
		return err
	}

	if ws.db == nil {
		log.Fatal("No database given to ", ws.Name())
	}

	conn, err := ws.db.Conn(ctx)
	if err != nil {
		return err
	}
	defer conn.Close()

	err = createTables(conn, ctx)
	if err != nil {
		return err
	}
	conn.Close()

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

func (ws WebServer) DBConn(ctx context.Context) (*sql.Conn, error) {
	return ws.db.Conn(ctx)
}
