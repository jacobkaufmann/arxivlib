package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"sync"
)

// A Datastore accesses the database
type Datastore struct {
	connectOnce sync.Once
	cfg         Config

	Database *sql.DB
}

// A Config holds configuration information about a datastore
type Config struct {
	Host         string
	Port         int
	Username     string
	Password     string
	DatabaseName string
	ModeSSL      string
}

// Connect connects to the MongoDB database specified by the Datastore Config
func (s *Datastore) Connect() {
	s.connectOnce.Do(func() {
		var err error

		connStr := fmt.Sprintf("postgres://%s:%s@%s:%d?sslmode=%s", s.cfg.Username, s.cfg.Password, s.cfg.Host, s.cfg.Port, s.cfg.ModeSSL)
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
func (s *Datastore) Ping() error {
	return s.Database.Ping()
}

// Transact executes query in a transaction, rolling back on failure and committing on success
func (s *Datastore) Transact(opt *sql.TxOptions, stmt *sql.Stmt, args ...interface{}) error {
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
