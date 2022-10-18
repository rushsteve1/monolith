package webserver

import (
	"rushsteve1.us/monolith/shared"
	ws "rushsteve1.us/monolith/webserver"
)

func main() {
	shared.MainHelper(
		&ws.WebServer{
			Config: shared.ConfigFromArgs(),
			Fcgi:   false},
		ws.Name)
}
