package swissarmybot

import (
	"context"
	"database/sql"
	"fmt"
	"io"
	"net/http"

	log "github.com/sirupsen/logrus"
	"rushsteve1.us/monolith/shared"
)

type SwissArmyBot struct {
	config shared.Config
	dbConn *sql.Conn
}

func New(ctx context.Context, cfg shared.Config, db *sql.DB) *SwissArmyBot {
	conn, err := db.Conn(ctx)
	if err != nil {
		log.Fatal(err)
	}
	return &SwissArmyBot{config: cfg, dbConn: conn}
}

func (sab *SwissArmyBot) Serve(ctx context.Context) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	err := createTables(sab.dbConn, ctx)
	if err != nil {
		return err
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/hello",
		func(w http.ResponseWriter, r *http.Request) {
			io.WriteString(w, "SAB is running\n")
		})

	defer sab.dbConn.Close()
	return shared.ServeHelper(mux, sab)
}

func (sab SwissArmyBot) Addr() string {
	return sab.config.SwissArmyBot.Addr
}

func (sab SwissArmyBot) Name() string {
	return "SwissArmyBot"
}

func (sab SwissArmyBot) UseFcgi() bool {
	return sab.config.UseFcgi
}

func (sab SwissArmyBot) String() string {
	return fmt.Sprintf("%s on %s", sab.Name(), sab.Addr())
}
