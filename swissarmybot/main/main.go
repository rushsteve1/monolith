package main

import (
	"rushsteve1.us/monolith/shared"
	sab "rushsteve1.us/monolith/swissarmybot"
)

func main() {
	shared.MainHelper(
		&sab.SwissArmyBot{
			Config: shared.ConfigFromArgs(),
			Fcgi:   false},
		sab.Name)
}
