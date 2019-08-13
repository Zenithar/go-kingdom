// Code generated by protoc-gen-defaults. DO NOT EDIT.

package userv1

import (
	"context"

	"github.com/bxcodec/faker"
)

// MockUserAPIServer is the mock implementation of the UserAPIServer. Use this to create mock services that
// return random data. Useful in UI Testing.
type MockUserAPIServer struct{}

// Create is mock implementation of the method Create
func (MockUserAPIServer) Create(context.Context, *CreateRequest) (*CreateResponse, error) {
	var res CreateResponse
	if err := faker.FakeData(&res); err != nil {
		return nil, err
	}
	return &res, nil
}

// Get is mock implementation of the method Get
func (MockUserAPIServer) Get(context.Context, *GetRequest) (*GetResponse, error) {
	var res GetResponse
	if err := faker.FakeData(&res); err != nil {
		return nil, err
	}
	return &res, nil
}

// Update is mock implementation of the method Update
func (MockUserAPIServer) Update(context.Context, *UpdateRequest) (*UpdateResponse, error) {
	var res UpdateResponse
	if err := faker.FakeData(&res); err != nil {
		return nil, err
	}
	return &res, nil
}

// Delete is mock implementation of the method Delete
func (MockUserAPIServer) Delete(context.Context, *DeleteRequest) (*DeleteResponse, error) {
	var res DeleteResponse
	if err := faker.FakeData(&res); err != nil {
		return nil, err
	}
	return &res, nil
}

// Search is mock implementation of the method Search
func (MockUserAPIServer) Search(context.Context, *SearchRequest) (*SearchResponse, error) {
	var res SearchResponse
	if err := faker.FakeData(&res); err != nil {
		return nil, err
	}
	return &res, nil
}

// Authenticate is mock implementation of the method Authenticate
func (MockUserAPIServer) Authenticate(context.Context, *AuthenticateRequest) (*AuthenticateResponse, error) {
	var res AuthenticateResponse
	if err := faker.FakeData(&res); err != nil {
		return nil, err
	}
	return &res, nil
}