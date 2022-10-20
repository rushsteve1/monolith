package main

import (
	"rushsteve1.us/monolith/shared"
	ws "rushsteve1.us/monolith/webserver"
)

func main() {
	serv := ws.WebServer{
		Config: shared.ConfigFromArgs(),
		Fcgi:   false}
	shared.MainHelper(&serv, serv.Name())
}
