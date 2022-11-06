package swissarmybot

import (
	"context"
	"database/sql"
	"fmt"
	"io"
	"net/http"

	"rushsteve1.us/monolith/shared"
)

type SwissArmyBotWeb struct {
	config shared.Config
	db     *sql.DB
}

func (sab *SwissArmyBotWeb) Serve(ctx context.Context) error {
	var err error
	setupOnce.Do(func() { err = setup(sab.db, ctx) })
	if err != nil {
		return err
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/",
		func(w http.ResponseWriter, r *http.Request) {
			io.WriteString(w, "SAB is running\n")
		})

	return shared.ServeHelper(mux, sab)
}

func (sab SwissArmyBotWeb) Addr() string {
	return sab.config.SwissArmyBot.Addr
}

func (sab SwissArmyBotWeb) Name() string {
	return "SwissArmyBot Web"
}

func (sab SwissArmyBotWeb) UseFcgi() bool {
	return sab.config.UseFcgi
}

func (sab SwissArmyBotWeb) String() string {
	return fmt.Sprintf("%s on %s", sab.Name(), sab.Addr())
}
