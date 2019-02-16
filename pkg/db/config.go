package db

// A Config holds configuration information about a datastore
type Config struct {
	Host         string
	Port         int
	Username     string
	Password     string
	DatabaseName string
}
