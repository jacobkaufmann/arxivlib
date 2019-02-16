package db

// A Datastore accesses a database
type Datastore interface {
	Connect()
	Ping() error
}
