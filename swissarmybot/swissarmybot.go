package swissarmybot

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"rushsteve1.us/monolith/shared"
)

const Name = "SwissArmyBot"

type SwissArmyBot struct {
	Config shared.Config
	Fcgi   bool
}

func (sab *SwissArmyBot) Serve(ctx context.Context) error {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello",
		func(w http.ResponseWriter, r *http.Request) {
			io.WriteString(w, "SAB is running\n")
		})

	return shared.ServeHelper(mux, sab)
}

func (sab *SwissArmyBot) Addr() string {
	return sab.Config.SwissArmyBot.Addr
}

func (sab *SwissArmyBot) Name() string {
	return Name
}

func (sab *SwissArmyBot) UseFcgi() bool {
	return sab.Fcgi
}

func (sab *SwissArmyBot) String() string {
	return fmt.Sprintf("%s on %s", sab.Name(), sab.Addr())
}
