package storage

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog"
	migrate "github.com/rubenv/sql-migrate"
)

func NewProvider(connstring, path string, log *zerolog.Logger) (*Provider, error) {
	p := Provider{
		log: log,
	}
	var err error

	p.conn, err = sql.Open("postgres", connstring)
	if err != nil {
		return nil, err
	}
	migrations := &migrate.FileMigrationSource{
		Dir: path,
	}

	n, err := migrate.Exec(p.conn, "postgres", migrations, migrate.Up)
	if err != nil {
		return nil, err
	}

	p.log.Info().Msgf("applied %d migrations to db", n)


	return &p, p.conn.Ping()
}
