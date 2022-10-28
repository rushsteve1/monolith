package main

import (
	"context"
	"database/sql"

	log "github.com/sirupsen/logrus"
	"github.com/thejerf/suture/v4"

	_ "modernc.org/sqlite"
	"rushsteve1.us/monolith/shared"
	sab "rushsteve1.us/monolith/swissarmybot"
	ws "rushsteve1.us/monolith/webserver"
)

type StoredService struct {
	Service shared.Service
	Token   suture.ServiceToken
}

var TopSup *suture.Supervisor
var ServiceMap map[string]StoredService

func main() {
	cfg := shared.ConfigFromArgs()

	TopSup = suture.NewSimple("overseer")

	var db *sql.DB
	if cfg.Database.UseSqlite {
		var err error
		db, err = sql.Open("sqlite", cfg.Database.String())
		if err != nil {
			log.Fatalf("Failed to connect to database: %v", err)
		}
	} else {
		log.Fatal("Only SQLite is supported right now")
	}

	ServiceMap = make(map[string]StoredService)
	services := []shared.Service{
		&Overseer{Config: cfg},
		&Cron{Config: cfg},
		&ws.WebServer{Config: cfg, Fcgi: cfg.Overseer.Fcgi, Database: db},
		&sab.SwissArmyBot{Config: cfg, Fcgi: cfg.Overseer.Fcgi, Database: db},
	}

	for _, serv := range services {
		token := TopSup.Add(serv)
		ServiceMap[serv.Name()] = StoredService{Service: serv, Token: token}
	}

	ctx, cancel := context.WithCancel(context.Background())
	TopSup.Serve(ctx)

	log.Warn("Top level supervisior exited unexpectedly")
	cancel()
}
