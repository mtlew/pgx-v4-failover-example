package main

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
)

type storage struct {
	conn *pgxpool.Pool
}

func newStorage(conn *pgxpool.Pool) *storage {
	return &storage{conn: conn}
}

type test struct {
	ID    string
	Title string
}

func (s *storage) first() (test, error) {
	var row test

	rows, _ := s.conn.Query(context.Background(), "SELECT id, title FROM test LIMIT 1")
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&row.ID, &row.Title)
		if err != nil {
			return row, err
		}
		break
	}

	return row, rows.Err()
}

func (s *storage) list() ([]test, error) {
	tests := make([]test, 0)

	rows, _ := s.conn.Query(context.Background(), "SELECT id, title FROM test")
	defer rows.Close()

	for rows.Next() {
		var row test
		err := rows.Scan(&row.ID, &row.Title)
		if err != nil {
			return tests, err
		}
		tests = append(tests, row)
	}

	return tests, rows.Err()
}
