package user

import (
	"fmt"

	"go.zenithar.org/kingdom/internal/models"
	userv1 "go.zenithar.org/kingdom/pkg/gen/go/kingdom/user/v1"
)

// FromEntity converts entity object to service object
func FromEntity(entity *models.User) *userv1.User {
	return &userv1.User{
		Urn: fmt.Sprintf("kndm:v1::user:%s:%s", entity.RealmID, entity.ID),
	}
}

// FromCollection returns a service object collection from entities
func FromCollection(entities []*models.User) []*userv1.User {
	res := make([]*userv1.User, len(entities))

	for i, entity := range entities {
		res[i] = FromEntity(entity)
	}

	return res
}
