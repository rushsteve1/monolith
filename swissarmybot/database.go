package swissarmybot

import (
	"context"
	"database/sql"
)

const createTablesSql = `
PRAGMA foreign_keys = ON;
PRAGMA journal_mode = WAL;

CREATE TABLE IF NOT EXISTS quotes (
       id INTEGER PRIMARY KEY NOT NULL,
       text TEXT NOT NULL,
       user_id INTEGER NOT NULL,
       user_name TEXT NOT NULL,
       author_id INTEGER NOT NULL,
       author_name TEXT NOT NULL,
       inserted_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);`

func createTables(db *sql.Conn, ctx context.Context) error {
	_, err := db.ExecContext(ctx, createTablesSql)
	if err != nil {
		return err
	}
	return nil
}
