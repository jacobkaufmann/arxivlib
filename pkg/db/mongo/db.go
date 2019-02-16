package mongo

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/mongodb/mongo-go-driver/mongo/readpref"
)

// A Datastore accesses the database
type Datastore struct {
	connectOnce sync.Once
	cfg         Config

	Client     *mongo.Client
	Collection *mongo.Collection
}

// A Config holds configuration information about a datastore
type Config struct {
	Host           string
	Port           int
	Username       string
	Password       string
	DatabaseName   string
	CollectionName string
}

// NewDatastore returns a new Datastore specified by a Config
func NewDatastore(cfg Config) Datastore {
	s := Datastore{cfg: cfg}
	return s
}

// Connect connects to the MongoDB database specified by the Datastore Config
func (s *Datastore) Connect() {
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

		s.Collection = s.Client.Database(s.cfg.DatabaseName).Collection(s.cfg.CollectionName)
	})
}

// Ping tests the connection to the mongo database
func (s *Datastore) Ping() error {
	ctx, cancel := context.WithTimeout(context.Background(), 2)
	defer cancel()
	err := s.Client.Ping(ctx, readpref.Primary())
	return err
}
