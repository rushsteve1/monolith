package main

import (
	"context"
	"database/sql"
	"time"

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
	log.Info("Starting Monolith Overseer")

	cfg := shared.ConfigFromArgs()

	log.SetLevel(log.Level(cfg.LogLevel))

	TopSup = suture.NewSimple("overseer")

	var db *sql.DB
	if cfg.Database.UseSqlite {
		var err error
		db, err = sql.Open("sqlite", cfg.Database.String())
		if err != nil {
			log.Fatalf("Failed to connect to database: %v", err)
		}
		db.SetMaxOpenConns(1)
		db.SetMaxIdleConns(8)
	} else {
		log.Fatal("Only SQLite is supported right now")
	}

	ctx, cancel := context.WithCancel(context.Background())

	pingCtx, pingCancel := context.WithTimeout(ctx, time.Duration(10_000))
	err := db.PingContext(pingCtx)
	if err != nil {
		log.Fatal(err)
	}
	defer pingCancel()

	log.Trace("Connected to Database")

	ServiceMap = make(map[string]StoredService)
	services := []shared.Service{
		&Overseer{Config: cfg},
		&Cron{Config: cfg},
		ws.New(ctx, cfg, db),
		sab.New(ctx, cfg, db),
	}

	log.Trace("Starting services...")

	for _, serv := range services {
		token := TopSup.Add(serv)
		ServiceMap[serv.Name()] = StoredService{Service: serv, Token: token}
	}

	log.Error(TopSup.Serve(ctx))

	cancel()
	log.Fatal("Top level supervisior exited unexpectedly, this is a hard-crash scenario")
}
