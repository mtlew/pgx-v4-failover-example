package main

import "log"

// https://github.com/jackc/pgx/pull/545
// https://www.postgresql.org/docs/current/libpq-connect.html + Ctrl+F = target_session_attrs
const connString = "postgres://postgres:example@localhost:5436,localhost:5437/postgres" // ?target_session_attrs=any

func main() {
	s := newStorage(newConn(connString))

	go func(s *storage) {
		p := newPing(s)
		p.run()
	}(s)

	log.Fatal(newHandler(s).init().run())
}
