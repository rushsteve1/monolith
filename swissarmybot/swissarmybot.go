package swissarmybot

import (
	"context"
	"database/sql"
	"fmt"
	"io"
	"net/http"

	"rushsteve1.us/monolith/shared"
)

type SwissArmyBot struct {
	Config   shared.Config
	Fcgi     bool
	Database *sql.DB
}

func (sab *SwissArmyBot) Serve(ctx context.Context) error {
	var err error

	err = createTables(sab.Database, ctx)
	if err != nil {
		return err
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/hello",
		func(w http.ResponseWriter, r *http.Request) {
			io.WriteString(w, "SAB is running\n")
		})

	return shared.ServeHelper(mux, sab)
}

func (sab SwissArmyBot) Addr() string {
	return sab.Config.SwissArmyBot.Addr
}

func (sab SwissArmyBot) Name() string {
	return "SwissArmyBot"
}

func (sab SwissArmyBot) UseFcgi() bool {
	return sab.Fcgi
}

func (sab SwissArmyBot) String() string {
	return fmt.Sprintf("%s on %s", sab.Name(), sab.Addr())
}
