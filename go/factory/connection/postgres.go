package connection

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

type PostgreSQL struct {
	db *sql.DB
}

func (p *PostgreSQL) Connect() error {
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		"postgres",
		"123456-78a",
		"127.0.0.1",
		"5432",
		"postgres",
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return err
	}

	p.db = db
	return nil
}

func (p *PostgreSQL) GetNow() (*time.Time, error) {
	t := &time.Time {}
	err := p.db.QueryRow("select CURRENT_DATE").Scan(t)
	if err != nil {
		fmt.Printf("Error al leer la fecha del servidor: %v", err)
		return nil, err
	}

	return t, nil
}

func (p *PostgreSQL) Close() error {
	return p.db.Close()
}
