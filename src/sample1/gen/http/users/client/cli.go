// Code generated by goa v3.2.6, DO NOT EDIT.
//
// Users HTTP client CLI support package
//
// Command:
// $ goa gen sample1/design

package client

import (
	"encoding/json"
	"fmt"
	users "sample1/gen/users"
	"unicode/utf8"

	goa "goa.design/goa/v3/pkg"
)

// BuildGetUserPayload builds the payload for the Users get user endpoint from
// CLI flags.
func BuildGetUserPayload(usersGetUserID string) (*users.GetUserPayload, error) {
	var id string
	{
		id = usersGetUserID
	}
	v := &users.GetUserPayload{}
	v.ID = id

	return v, nil
}

// BuildCreateUserPayload builds the payload for the Users create user endpoint
// from CLI flags.
func BuildCreateUserPayload(usersCreateUserBody string) (*users.CreateUserPayload, error) {
	var err error
	var body CreateUserRequestBody
	{
		err = json.Unmarshal([]byte(usersCreateUserBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"email\": \"Aperiam eos.\",\n      \"id\": \"XRQ85mtXnINISH25zfM0m5RlC6L2\",\n      \"name\": \"Ipsum laborum numquam.\"\n   }'")
		}
		if utf8.RuneCountInString(body.ID) < 1 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.id", body.ID, utf8.RuneCountInString(body.ID), 1, true))
		}
		if utf8.RuneCountInString(body.ID) > 28 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.id", body.ID, utf8.RuneCountInString(body.ID), 28, false))
		}
		if err != nil {
			return nil, err
		}
	}
	v := &users.CreateUserPayload{
		ID:    body.ID,
		Name:  body.Name,
		Email: body.Email,
	}

	return v, nil
}

// BuildUpdateUserPayload builds the payload for the Users update user endpoint
// from CLI flags.
func BuildUpdateUserPayload(usersUpdateUserBody string, usersUpdateUserID string) (*users.UpdateUserPayload, error) {
	var err error
	var body UpdateUserRequestBody
	{
		err = json.Unmarshal([]byte(usersUpdateUserBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"email\": \"Officia perspiciatis occaecati laboriosam natus autem laborum.\",\n      \"name\": \"Sit sequi et rerum.\"\n   }'")
		}
	}
	var id string
	{
		id = usersUpdateUserID
		if utf8.RuneCountInString(id) < 1 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("id", id, utf8.RuneCountInString(id), 1, true))
		}
		if utf8.RuneCountInString(id) > 28 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("id", id, utf8.RuneCountInString(id), 28, false))
		}
		if err != nil {
			return nil, err
		}
	}
	v := &users.UpdateUserPayload{
		Name:  body.Name,
		Email: body.Email,
	}
	v.ID = id

	return v, nil
}

// BuildDeleteUserPayload builds the payload for the Users delete user endpoint
// from CLI flags.
func BuildDeleteUserPayload(usersDeleteUserID string) (*users.DeleteUserPayload, error) {
	var id string
	{
		id = usersDeleteUserID
	}
	v := &users.DeleteUserPayload{}
	v.ID = id

	return v, nil
}
