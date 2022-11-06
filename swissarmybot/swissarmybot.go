package swissarmybot

import (
	"context"
	"database/sql"
	"sync"

	"github.com/thejerf/suture/v4"
	"rushsteve1.us/monolith/shared"
)

func NewSupervisor(ctx context.Context, cfg shared.Config, db *sql.DB) (*suture.Supervisor, map[string]shared.StoredService) {
	sabWeb := SwissArmyBotWeb{config: cfg, db: db}
	sabCord := SwissArmyBotDiscord{config: cfg, db: db}

	sabSup := suture.NewSimple("SwissArmyBot")
	servMap := make(map[string]shared.StoredService, 2)

	servMap[sabWeb.Name()] = shared.StoredService{Service: &sabWeb, Token: sabSup.Add(&sabWeb)}
	servMap[sabCord.Name()] = shared.StoredService{Service: &sabCord, Token: sabSup.Add(&sabCord)}

	return sabSup, servMap
}

var setupOnce sync.Once

func setup(db *sql.DB, ctx context.Context) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	conn, err := db.Conn(ctx)
	if err != nil {
		return err
	}
	defer conn.Close()

	err = createTables(conn, ctx)
	if err != nil {
		return err
	}
	conn.Close()
	return nil
}
