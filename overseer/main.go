package main

import (
	"context"
	"net"
	"net/http"
	"net/http/fcgi"
	"net/rpc"

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

	if cfg.Overseer.Rpc {
		rpcobj := new(OverseerRpc)
		rpc.RegisterName("Overseer", rpcobj)
		rpc.HandleHTTP()

		listener, err := net.Listen("tcp", cfg.Overseer.Addr)
		if err != nil {
			log.Fatal("Error starting Overseer RPC: ", err)
		}

		log.Info("Overseer RPC server started on ", cfg.Overseer.Addr)
		if cfg.Overseer.Fcgi {
			go fcgi.Serve(listener, nil)
		} else {
			go http.Serve(listener, nil)
		}
	}

	ServiceMap[ws.Name] = TopSup.Add(&ws.WebServer{Config: cfg, Fcgi: cfg.Overseer.Fcgi})
	ServiceMap[sab.Name] = TopSup.Add(&sab.SwissArmyBot{Config: cfg, Fcgi: cfg.Overseer.Fcgi})

	ctx, cancel := context.WithCancel(context.Background())
	TopSup.Serve(ctx)

	log.Warn("Top level supervisior exited unexpectedly")
	cancel()
}
