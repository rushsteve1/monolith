package main

import (
	"context"

	"rushsteve1.us/monolith/shared"
	ws "rushsteve1.us/monolith/webserver"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	serv := ws.New(ctx, shared.ConfigFromArgs(), nil)
	shared.MainHelper(serv, serv.Name())
}
