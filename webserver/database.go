package webserver

import (
	"context"
	"database/sql"
	"time"
)

type BlogPost struct {
	ID         int64
	Title      string
	Body       string
	Published  bool
	InsertedAt time.Time
	UpdatedAt  time.Time
}

const createTablesSql = `
PRAGMA foreign_keys = ON;
PRAGMA journal_mode = WAL;

CREATE TABLE IF NOT EXISTS blog (
    id INTEGER PRIMARY KEY NOT NULL,
	title TEXT NOT NULL,
    body TEXT NOT NULL,
    published BOOLEAN NOT NULL DEFAULT false,
    inserted_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS kayvee (
    key VARCHAR(64) PRIMARY KEY NOT NULL,
    value JSONB
);`

func createTables(db *sql.DB, ctx context.Context) error {
	_, err := db.ExecContext(ctx, createTablesSql)
	if err != nil {
		return err
	}
	return nil
}

func ListPosts(db *sql.DB, ctx context.Context) ([]BlogPost, error) {
	rows, err := db.QueryContext(ctx, "SELECT * FROM blog;")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []BlogPost
	for rows.Next() {
		var post BlogPost
		err = rows.Scan(&post.ID, &post.Title, &post.Body, &post.Published, &post.InsertedAt, &post.UpdatedAt)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return posts, nil
}

func GetPost(db *sql.DB, ctx context.Context, id int64) (BlogPost, error) {
	var post BlogPost
	row := db.QueryRowContext(ctx, "SELECT * FROM blog WHERE id = ?;", id)
	if row.Err() != nil {
		return post, row.Err()
	}

	err := row.Scan(&post.ID, &post.Title, &post.Body, &post.Published, &post.InsertedAt, &post.UpdatedAt)
	return post, err
}

func InsertPost(db *sql.DB, ctx context.Context, title string, body string) error {
	_, err := db.ExecContext(ctx, "INSERT INTO blog (title, body) VALUES (?, ?);", title, body)
	return err
}

func ChangePublished(db *sql.DB, ctx context.Context, id int64, published bool) error {
	_, err := db.ExecContext(ctx, "UPDATE blog SET published = ?, updated_at = ? WHERE id = ?;", published, time.Now(), id)
	return err
}

func UpdatePost(db *sql.DB, ctx context.Context, id int64, title string, body string) error {
	_, err := db.ExecContext(ctx, "UPDATE blog SET title = ?, body = ?, updated_at = ? WHERE id = ?;", title, body, time.Now(), id)
	return err
}

func DeletePost(db *sql.DB, ctx context.Context, id int64) error {
	_, err := db.ExecContext(ctx, "DEELETE FROM blog WHERE id = ?;", id)
	return err
}
