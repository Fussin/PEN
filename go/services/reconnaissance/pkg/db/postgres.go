package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"time"
)

type Postgres struct {
	db *sql.DB
}

func NewPostgres(connStr string) (*Postgres, error) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	// Retry connecting to the database
	for i := 0; i < 5; i++ {
		err = db.Ping()
		if err == nil {
			break
		}
		time.Sleep(2 * time.Second)
	}
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	return &Postgres{db: db}, nil
}

func (p *Postgres) SaveResult(result string) error {
	_, err := p.db.Exec("INSERT INTO results (result) VALUES ($1)", result)
	return err
}
