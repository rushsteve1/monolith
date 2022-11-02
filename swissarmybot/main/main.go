package main

import (
	"context"

	"rushsteve1.us/monolith/shared"
	sab "rushsteve1.us/monolith/swissarmybot"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	serv := sab.New(ctx, shared.ConfigFromArgs(), nil)
	shared.MainHelper(serv, serv.Name())
}
