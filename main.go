package main

import "log"

// https://github.com/jackc/pgx/pull/545
// https://www.postgresql.org/docs/current/libpq-connect.html + Ctrl+F = target_session_attrs
const connString = "postgres://postgres:example@localhost:5436,localhost:5437/postgres?pool_max_conns=1000" // ?target_session_attrs=any
const goroutines = 5

func main() {
	s := newStorage(newConn(connString))

	for i := 0; i < goroutines; i++ {
		go func(s *storage) {
			p := newPing(s)
			p.run()
		}(s)
	}

	log.Fatal(newHandler(s).init().run())
}
