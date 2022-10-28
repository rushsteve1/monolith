package swissarmybot

import (
	"context"
	"database/sql"
	"embed"
)

//go:embed sql
var sqlFiles embed.FS

func createTables(db *sql.DB, ctx context.Context) error {
	sqlText, err := sqlFiles.ReadFile("sql/create_tables.sql")
	if err != nil {
		return err
	}
	_, err = db.ExecContext(ctx, string(sqlText))
	if err != nil {
		return err
	}
	return nil
}
