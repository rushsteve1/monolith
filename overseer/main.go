package main

import (
	"context"

	log "github.com/sirupsen/logrus"
	"github.com/thejerf/suture/v4"

	"rushsteve1.us/monolith/shared"
	sab "rushsteve1.us/monolith/swissarmybot"
	ws "rushsteve1.us/monolith/webserver"
)

var TopSup *suture.Supervisor
var ServiceMap = map[string]suture.ServiceToken{}

func main() {
	cfg := shared.ConfigFromArgs()

	TopSup = suture.NewSimple("overseer")

	services := []shared.Service{
		&Overseer{Config: cfg},
		&Cron{Config: cfg},
		&ws.WebServer{Config: cfg, Fcgi: cfg.Overseer.Fcgi},
		&sab.SwissArmyBot{Config: cfg, Fcgi: cfg.Overseer.Fcgi},
	}

	for _, serv := range services {
		ServiceMap[serv.Name()] = TopSup.Add(serv)
	}

	ctx, cancel := context.WithCancel(context.Background())
	TopSup.Serve(ctx)

	log.Warn("Top level supervisior exited unexpectedly")
	cancel()
}
