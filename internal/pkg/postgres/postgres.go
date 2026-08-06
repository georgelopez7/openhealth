package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
)

type Postgres struct {
	DB *sqlx.DB
}

func NewPostgresDB(uri string) *Postgres {
	db, err := sqlx.Connect("postgres", uri)
	if err != nil {
		log.Fatalln(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal("failed to ping postgres: ", err)
	}

	return &Postgres{
		DB: db,
	}
}

// Migrate - runs migrations
func (p *Postgres) Migrate(path string) {
	if err := goose.SetDialect("postgres"); err != nil {
		log.Fatal("failed to set goose dialect: ", err)
	}

	if err := goose.Up(p.DB.DB, path); err != nil {
		log.Fatal("failed to run migrations: ", err)
	}
}

// TxManager - manager to handle transactions
type TxManager struct {
	db *sqlx.DB
}

func NewTxManager(db *sqlx.DB) *TxManager {
	return &TxManager{db: db}
}

type txKey struct{}

func (tm *TxManager) WithTransaction(ctx context.Context, fn func(ctx context.Context) error) error {
	tx, err := tm.db.BeginTxx(ctx, &sql.TxOptions{})
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}

	ctx = context.WithValue(ctx, txKey{}, tx)

	defer func() {
		if p := recover(); p != nil {
			_ = tx.Rollback()
			panic(p)
		}
	}()

	err = fn(ctx)
	if err != nil {
		txErr := tx.Rollback()
		if txErr != nil {
			return fmt.Errorf("failed to rollback transaction: %w", txErr)
		}

		return err
	}

	return tx.Commit()
}

// GetTxOrDB - gets the transaction client or the database client from the context
func GetTxOrDB(ctx context.Context, db *sqlx.DB) sqlx.ExtContext {
	if tx, ok := ctx.Value(txKey{}).(*sqlx.Tx); ok {
		return tx
	}

	return db
}
