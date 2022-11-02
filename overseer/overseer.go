package main

import (
	"context"
	_ "expvar"
	"fmt"
	"net"
	"net/http"
	"net/http/fcgi"
	_ "net/http/pprof"
	"net/rpc"

	log "github.com/sirupsen/logrus"

	"rushsteve1.us/monolith/shared"
)

const Name string = "Overseer"

type Overseer struct {
	Config shared.Config
}

func (ov *Overseer) Serve(ctx context.Context) error {
	if ov.Config.Overseer.Rpc {
		rpcobj := new(OverseerRpc)
		rpcobj.Config = ov.Config
		rpc.RegisterName("Overseer", rpcobj)
		rpc.HandleHTTP()
	}

	if ov.Config.Overseer.Debug || ov.Config.Overseer.Rpc {
		// This server uses the DefualtServMux
		listener, err := net.Listen("tcp", ov.Config.Overseer.Addr)
		if err != nil {
			log.Fatal("Error starting Overseer RPC: ", err)
		}

		log.Info("Overseer server started on ", ov.Config.Overseer.Addr)
		if ov.UseFcgi() {
			return fcgi.Serve(listener, nil)
		} else {
			return http.Serve(listener, nil)
		}
	}

	return nil
}

func (ov Overseer) Addr() string {
	return ov.Config.Overseer.Addr
}

func (ov Overseer) Name() string {
	return Name
}

func (ov Overseer) UseFcgi() bool {
	return ov.Config.UseFcgi
}

func (ov Overseer) String() string {
	return fmt.Sprintf("%s on %s", ov.Name(), ov.Addr())
}
