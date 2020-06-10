package mysql_repository

import (
	"context"
	"database/sql"
)

func queryRow(ctx context.Context, conn *sql.DB, q string, args ...interface{}) (*sql.Row, error) {
	stmt, err := conn.Prepare(q)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	return stmt.QueryRowContext(ctx, args...), nil
}

func query(ctx context.Context, conn *sql.DB, q string, args ...interface{}) (*sql.Rows, error) {
	stmt, err := conn.Prepare(q)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	return stmt.QueryContext(ctx, args...)
}
