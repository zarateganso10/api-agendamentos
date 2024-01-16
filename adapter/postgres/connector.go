package postgres

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type PostgresDatabase struct {
	URL string
}

func NewPostgresDatabase(url string) *PostgresDatabase {
	return &PostgresDatabase{
		URL: url,
	}
}

func (pg *PostgresDatabase) OpenConnection() (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", pg.URL)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	return db, err
}
