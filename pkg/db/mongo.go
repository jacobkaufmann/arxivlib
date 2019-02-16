package db

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/mongodb/mongo-go-driver/mongo/readpref"
)

// A MongoDatastore accesses the database
type MongoDatastore struct {
	connectOnce sync.Once
	cfg         Config

	Client   *mongo.Client
	Database *mongo.Database
}

// NewMongoDatastore returns a new MongoDatastore specified by a Config
func NewMongoDatastore(cfg Config) MongoDatastore {
	s := MongoDatastore{cfg: cfg}
	return s
}

// Connect connects to the MongoDB database specified by the Datastore Config
func (s *MongoDatastore) Connect() {
	s.connectOnce.Do(func() {
		var err error

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		connStr := fmt.Sprintf("mongodb://%s:%s@%s:%d", s.cfg.Username, s.cfg.Password, s.cfg.Host, s.cfg.Port)
		s.Client, err = mongo.Connect(ctx, connStr)
		if err != nil {
			log.Fatal(err)
		}

		err = s.Ping()
		if err != nil {
			log.Fatal(err)
		}

		s.Database = s.Client.Database(s.cfg.DatabaseName)
	})
}

// Ping tests the connection to the mongo database
func (s *MongoDatastore) Ping() error {
	ctx, cancel := context.WithTimeout(context.Background(), 2)
	defer cancel()
	err := s.Client.Ping(ctx, readpref.Primary())
	return err
}
