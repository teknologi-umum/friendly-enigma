package migration

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
)

func Migrate(db *sql.DB, ctx context.Context) error {
	err := tableFamily(db, ctx)
	if err != nil {
		return err
	}

	err = tableRefrigerator(db, ctx)
	if err != nil {
		return err
	}

	err = insertData(db, ctx)
	if err != nil {
		return err
	}

	return nil
}

func tableFamily(db *sql.DB, ctx context.Context) error {
	c, err := db.Conn(ctx)
	if err != nil {
		return err
	}
	defer c.Close()

	_, err = c.ExecContext(
		ctx,
		`CREATE TABLE IF NOT EXISTS family (
			id VARCHAR(36) PRIMARY KEY,
			name VARCHAR(255) UNIQUE NOT NULL,
			permission BLOB NOT NULL
		);`,
	)
	if err != nil {
		return err
	}

	return nil
}

func tableRefrigerator(db *sql.DB, ctx context.Context) error {
	c, err := db.Conn(ctx)
	if err != nil {
		return err
	}
	defer c.Close()

	_, err = c.ExecContext(
		ctx,
		`CREATE TABLE IF NOT EXISTS refrigerator (
			id VARCHAR(36) PRIMARY KEY,
			name VARCHAR(255) NOT NULL,
			quantity INT NOT NULL,
			updated_at TIMESTAMP,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP 
		);`,
	)
	if err != nil {
		return err
	}

	return nil
}

func insertData(db *sql.DB, ctx context.Context) error {
	c, err := db.Conn(ctx)
	if err != nil {
		return err
	}
	defer c.Close()

	t, err := c.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return err
	}

	_, err = t.ExecContext(
		ctx,
		`INSERT INTO family
			(id, name, permission)
			VALUES
			(?, ?, ?),
			(?, ?, ?),
			(?, ?, ?),
			(?, ?, ?),
			(?, ?, ?),`,
		uuid.NewString(),
		"father",
		0b1111,
		uuid.NewString(),
		"mother",
		0b1111,
		uuid.NewString(),
		"brother",
		0b1110,
		uuid.NewString(),
		"sister",
		0b1110,
		uuid.NewString(),
		"grandma",
		0b0000,
	)
	if err != nil {
		t.Rollback()
		return err
	}

	_, err = t.ExecContext(
		ctx,
		`INSERT INTO refrigerator
			(id, name, quantity, updated_at)
			VALUES
			(?, ?, ?, ?);`,
		uuid.NewString(),
		"Chicken Chowder",
		8,
		time.Now(),
	)
	if err != nil {
		t.Rollback()
		return err
	}

	err = t.Commit()
	if err != nil {
		return err
	}

	return nil
}
