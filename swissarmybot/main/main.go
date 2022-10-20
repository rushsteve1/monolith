package main

import (
	"rushsteve1.us/monolith/shared"
	sab "rushsteve1.us/monolith/swissarmybot"
)

func main() {
	serv := sab.SwissArmyBot{
		Config: shared.ConfigFromArgs(),
		Fcgi:   false}
	shared.MainHelper(&serv, serv.Name())
}
