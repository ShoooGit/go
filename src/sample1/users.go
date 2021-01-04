package goa3sample

import (
	"context"
	"database/sql"
	"log"
	users "sample1/gen/users"
)

// Users service example implementation.
// The example methods log the requests and return zero values.
type userssrvc struct {
	logger *log.Logger
	DB     *sql.DB
}

// NewUsers returns the Users service implementation.
func NewUsers(logger *log.Logger, db *sql.DB) users.Service {
	return &userssrvc{logger, db}
}

// List all stored users
func (s *userssrvc) ListUser(ctx context.Context) (res users.Goa3SampleUserCollection, err error) {
	s.logger.Print("users.list user")
	return
}

// Show user by ID
func (s *userssrvc) GetUser(ctx context.Context, p *users.GetUserPayload) (res *users.Goa3SampleUser, err error) {
	res = &users.Goa3SampleUser{}
	s.logger.Print("users.get user")
	var id, name, email string
	err = s.DB.QueryRow("SELECT id, name, email FROM users WHERE id=?", p.ID).Scan(&id, &name, &email)
	if err != nil {
		return
	}
	res = &users.Goa3SampleUser{ID: id, Name: name, Email: email}
	return
}

// Add new user and return its ID.
func (s *userssrvc) CreateUser(ctx context.Context, p *users.CreateUserPayload) (res string, err error) {
	s.logger.Print("users.create user")
	return
}

// Update user item.
func (s *userssrvc) UpdateUser(ctx context.Context, p *users.UpdateUserPayload) (res *users.Goa3SampleUser, err error) {
	res = &users.Goa3SampleUser{}
	s.logger.Print("users.update user")
	return
}

// Delete user by id.
func (s *userssrvc) DeleteUser(ctx context.Context, p *users.DeleteUserPayload) (err error) {
	s.logger.Print("users.delete user")
	return
}
