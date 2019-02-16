package arxivlib

import (
	"errors"

	"github.com/mongodb/mongo-go-driver/bson/primitive"
)

// A User is a user on arxivlib
type User struct {
	ID       primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Username string             `json:"username" form:"username"`
	Passwd   string             `json:"passwd,omitempty" form:"passwd,omitempty"`
}

// UsersService interacts with the user-related endpoints in arxivlib's API
type UsersService interface {
	// Get a user
	Get(id primitive.ObjectID) (*User, error)

	// Authenticate a user
	Authenticate(username, passwd string) (*User, error)

	// List users
	List(opt *UserListOptions) ([]*User, error)

	// Create a user
	Create(u *User) (created bool, err error)

	// Delete a user
	Delete(id primitive.ObjectID) (deleted bool, err error)
}

var (
	// ErrUserNotFound is a failure to retrieve a specified user
	ErrUserNotFound = errors.New("user not found")
)

// A UserListOptions represents a search filter for listing users
type UserListOptions struct {
	Username string `json:"username,omitempty" bson:"username,omitempty" form:"username,omitempty"`
}
