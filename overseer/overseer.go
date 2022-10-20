package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"net/http/fcgi"
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
		rpc.RegisterName("Overseer", rpcobj)
		rpc.HandleHTTP()

		listener, err := net.Listen("tcp", ov.Config.Overseer.Addr)
		if err != nil {
			log.Fatal("Error starting Overseer RPC: ", err)
		}

		handler := shared.LogWrapper(http.DefaultServeMux, ov)

		log.Info("Overseer RPC server started on ", ov.Config.Overseer.Addr)
		if ov.Config.Overseer.Fcgi {
			return fcgi.Serve(listener, handler)
		} else {
			return http.Serve(listener, handler)
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
	return ov.Config.Overseer.Fcgi
}

func (ov Overseer) String() string {
	return fmt.Sprintf("%s on %s", ov.Name(), ov.Addr())
}
