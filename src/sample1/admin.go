package goa3sample

import (
	"context"
	"log"
	admin "sample1/gen/admin"
)

// Admin service example implementation.
// The example methods log the requests and return zero values.
type adminsrvc struct {
	logger *log.Logger
}

// NewAdmin returns the Admin service implementation.
func NewAdmin(logger *log.Logger) admin.Service {
	return &adminsrvc{logger}
}

// Number of users
func (s *adminsrvc) UserNumber(ctx context.Context) (res *admin.Goa3SampleAdminUserNumber, err error) {
	res = &admin.Goa3SampleAdminUserNumber{}
	s.logger.Print("admin.user_number")
	return
}

// List all stored users
func (s *adminsrvc) AdminListUser(ctx context.Context) (res admin.Goa3SampleAdminUserCollection, err error) {
	s.logger.Print("admin.admin list user")
	return
}

// Show user by ID
func (s *adminsrvc) AdminGetUser(ctx context.Context, p *admin.GetUserPayload) (res *admin.Goa3SampleAdminUser, err error) {
	res = &admin.Goa3SampleAdminUser{}
	s.logger.Print("admin.admin get user")
	return
}

// Add new user and return its ID.
func (s *adminsrvc) AdminCreateUser(ctx context.Context, p *admin.CreateUserPayload) (res string, err error) {
	s.logger.Print("admin.admin create user")
	return
}

// Update user item.
func (s *adminsrvc) AdminUpdateUser(ctx context.Context, p *admin.UpdateUserPayload) (res *admin.Goa3SampleAdminUser, err error) {
	res = &admin.Goa3SampleAdminUser{}
	s.logger.Print("admin.admin update user")
	return
}

// Delete user by id.
func (s *adminsrvc) AdminDeleteUser(ctx context.Context, p *admin.DeleteUserPayload) (err error) {
	s.logger.Print("admin.admin delete user")
	return
}
