package driver

import (
	"database/sql"
	"time"

	_ "github.com/lib/pq"
)

type DB struct {
	SQL *sql.DB
}

const (
	maxOpenDBConn = 10
	maxIdleDBConn = 5
	maxLifeDBTime = 5 * time.Minute
)

func ConnectDB(dbName, dsn string) (*DB, error) {
	conn, err := NewDatabase(dbName, dsn)
	if err != nil {
		return nil, err
	}
	conn.SetMaxOpenConns(maxOpenDBConn)
	conn.SetMaxIdleConns(maxIdleDBConn)
	conn.SetConnMaxLifetime(maxLifeDBTime)
	if err := TestDB(conn); err != nil {
		return nil, err
	}
	return &DB{
		SQL: conn,
	}, nil
}

func NewDatabase(dbName, dsn string) (*sql.DB, error) {
	db, err := sql.Open(dbName, dsn)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}

func TestDB(db *sql.DB) error {
	return db.Ping()
}
