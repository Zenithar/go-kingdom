package realm

import (
	"go.zenithar.org/kingdom/internal/models"
	realmv1 "go.zenithar.org/kingdom/pkg/gen/go/kingdom/realm/v1"
)

// FromEntity converts entity object to service object
func FromEntity(entity *models.Realm) *realmv1.Realm {
	return &realmv1.Realm{
		Id:    entity.ID,
		Label: entity.Label,
	}
}

// FromCollection returns a service object collection from entities
func FromCollection(entities []*models.Realm) []*realmv1.Realm {
	res := make([]*realmv1.Realm, len(entities))

	for i, entity := range entities {
		res[i] = FromEntity(entity)
	}

	return res
}
