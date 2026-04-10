package postgres

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/pgconn"
)

type Repository struct {
	pool *pgxpool.Pool
}

func New(pool *pgxpool.Pool) *Repository {
	return &Repository{pool: pool}
}

type DB interface {
	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
	Exec(ctx context.Context, sql string, args ...any) (pgconn.CommandTag, error)
}

func (r *Repository) getDB(tx any) DB {
    if tx != nil {
		pgxTx, ok := tx.(pgx.Tx)
        if ok {
            return pgxTx
        }
    }
    return r.pool
}

func (r *Repository) Begin(ctx context.Context) (any, error) {
    tx, err := r.pool.Begin(ctx)
    if err != nil {
        return nil, err
    }
    return tx, nil
}

func (r *Repository) Commit(ctx context.Context, tx any) error {
    pgxTx, ok := tx.(pgx.Tx)
    if !ok {
        return errors.New("invalid transaction type")
    }
    return pgxTx.Commit(ctx)
}

func (r *Repository) Rollback(ctx context.Context, tx any) error {
	pgxTx, ok := tx.(pgx.Tx)
    if !ok {
        return errors.New("invalid transaction type")
    }
    return pgxTx.Rollback(ctx)
}