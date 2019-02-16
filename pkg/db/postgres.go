package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"sync"
)

// A PostgresDatastore accesses a postgres database
type PostgresDatastore struct {
	connectOnce sync.Once
	cfg         Config
	modeSSL     string

	Database *sql.DB
}

// Connect connects to the MongoDB database specified by the Datastore Config
func (s *PostgresDatastore) Connect() {
	s.connectOnce.Do(func() {
		var err error

		connStr := fmt.Sprintf("postgres://%s:%s@%s:%d?sslmode=%s", s.cfg.Username, s.cfg.Password, s.cfg.Host, s.cfg.Port, s.modeSSL)
		s.Database, err = sql.Open("postgres", connStr)
		if err != nil {
			log.Fatal(err)
		}

		err = s.Ping()
		if err != nil {
			log.Fatal(err)
		}
	})
}

// Ping tests the connection to the mongo database
func (s *PostgresDatastore) Ping() error {
	return s.Database.Ping()
}

// Transact executes query in a transaction, rolling back on failure and committing on success
func (s *PostgresDatastore) Transact(opt *sql.TxOptions, stmt *sql.Stmt, args ...interface{}) error {
	ctx := context.Background()
	tx, err := s.Database.BeginTx(ctx, opt)
	if err != nil {
		log.Fatal(err)
	}

	_, execErr := tx.Stmt(stmt).Exec(args)
	if execErr != nil {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			log.Printf("Could not roll back: %v\n", rollbackErr)
		}
		log.Fatal(execErr)
	}
	if err := tx.Commit(); err != nil {
		log.Fatal(err)
	}
	return nil
}
