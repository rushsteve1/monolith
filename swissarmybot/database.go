package swissarmybot

import (
	"context"
	"database/sql"
	"time"
)

type PartialQuote struct {
	Text       string
	UserId     int
	UserName   string
	AuthorId   int
	AuthorName string
	ServerID   int
}

type Quote struct {
	ID int
	PartialQuote
	InsertedAt time.Time
}

const createTablesSql = `
PRAGMA foreign_keys = ON;
PRAGMA journal_mode = WAL;

CREATE TABLE IF NOT EXISTS quotes (
	id          INTEGER PRIMARY KEY NOT NULL,
	text        TEXT NOT NULL,
	user_id     INTEGER NOT NULL,
	user_name   TEXT NOT NULL,
	author_id   INTEGER NOT NULL,
	author_name TEXT NOT NULL,
	server_id   INTEGER NOT NULL,
	inserted_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);`

func createTables(conn *sql.Conn, ctx context.Context) error {
	_, err := conn.ExecContext(ctx, createTablesSql)
	if err != nil {
		return err
	}
	return nil
}

func ListQuotes(conn *sql.Conn, ctx context.Context) ([]Quote, error) {
	rows, err := conn.QueryContext(ctx, "SELECT * FROM quotes;")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var quotes []Quote
	for rows.Next() {
		var quote Quote
		err = rows.Scan(&quote.ID, &quote.Text, &quote.UserId, &quote.UserName, &quote.AuthorId, &quote.AuthorName, &quote.ServerID, &quote.InsertedAt)
		if err != nil {
			return nil, err
		}
		quotes = append(quotes, quote)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return quotes, nil
}

func GetQuote(conn *sql.Conn, ctx context.Context, id int) (Quote, error) {
	var quote Quote
	row := conn.QueryRowContext(ctx, "SELECT * FROM quotes WHERE id = ?;", id)
	if row.Err() != nil {
		return quote, row.Err()
	}

	err := row.Scan(&quote.ID, &quote.Text, &quote.UserId, &quote.UserName, &quote.AuthorId, &quote.AuthorName, &quote.ServerID, &quote.InsertedAt)
	return quote, err
}

func InsertQuote(conn *sql.Conn, ctx context.Context, quote PartialQuote) error {
	_, err := conn.ExecContext(ctx,
		"INSERT INTO quotes (text, user_id, user_name, author_id, author_name, server_id) VALUES (?, ?, ?, ?, ?, ?)",
		quote.Text, quote.UserId, quote.UserName, quote.AuthorId, quote.AuthorName, quote.ServerID)
	return err
}

func DeleteQuote(conn *sql.Conn, ctx context.Context, id int64) error {
	_, err := conn.ExecContext(ctx, "DELETE FROM quotes WHERE id = ?;", id)
	return err
}

func RandomQuoteFromServer(conn *sql.Conn, ctx context.Context, serverId int) (Quote, error) {
	var quote Quote
	row := conn.QueryRowContext(ctx, "SELECT * FROM quotes WHERE server_id = ? ORDER BY RAND() LIMIT 1;", serverId)
	if row.Err() != nil {
		return quote, row.Err()
	}

	err := row.Scan(&quote.ID, &quote.Text, &quote.UserId, &quote.UserName, &quote.AuthorId, &quote.AuthorName, &quote.ServerID, &quote.InsertedAt)
	return quote, err
}
