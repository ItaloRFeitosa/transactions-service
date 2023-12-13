package database

import (
	"context"

	"github.com/jmoiron/sqlx"

	_ "github.com/lib/pq"
)

func Init(databaseDSN string) (*sqlx.DB, error) {
	db, err := sqlx.Connect("postgres", databaseDSN)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func Close(db *sqlx.DB) error {
	return db.DB.Close()
}

type sqlxDB interface {
	PrepareNamedContext(ctx context.Context, query string) (*sqlx.NamedStmt, error)
}

func insertReturningID(ctx context.Context, db sqlxDB, query string, arg any) (int, error) {
	var id int

	stmt, err := db.PrepareNamedContext(ctx, query)
	if err != nil {
		return id, err
	}

	row := stmt.QueryRow(arg)
	if err := row.Err(); err != nil {
		return id, err
	}

	err = row.Scan(&id)
	if err != nil {
		return id, err
	}

	return id, nil
}
