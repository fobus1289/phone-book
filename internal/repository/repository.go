package repository

import (
	"context"
	"database/sql"
	"log"
)

type DB interface {
	QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error)
	QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row
	ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error)
}

type repository[T any] struct {
	db  DB
	ctx context.Context
}

type scannerFn[T any] func(*T, func(dest ...any) error) error

func NewRepositroy[T any](db DB, ctx context.Context) *repository[T] {
	if ctx == nil {
		ctx = context.Background()
	}

	return &repository[T]{
		db:  db,
		ctx: ctx,
	}
}

func (*repository[T]) WithTx(tx DB, ctx context.Context) *repository[T] {

	if ctx == nil {
		ctx = context.Background()
	}

	return &repository[T]{
		db:  tx,
		ctx: ctx,
	}
}

func (r *repository[T]) Find(query string, scan scannerFn[T], args ...any) ([]T, error) {

	rows, err := r.db.QueryContext(r.ctx, query, args...)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var items []T

	for rows.Next() {

		var item T

		if err := scan(&item, rows.Scan); err != nil {
			return nil, err
		}

		items = append(items, item)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return items, nil
}

func (r *repository[T]) FindOne(query string, scan scannerFn[T], args ...any) (*T, error) {

	row := r.db.QueryRowContext(r.ctx, query, args...)

	var item T

	log.Println(row.Err())

	if err := scan(&item, row.Scan); err != nil {
		return nil, err
	}

	return &item, nil
}
