package models

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"

	"go.zenithar.org/kingdom/internal/helpers"
)

// Realm represents an object group
type Realm struct {
	ID           string    `json:"id" db:"realm_id"`
	Label        string    `json:"label" db:"label"`
	CreationDate time.Time `json:"creation_date" db:"creation_date"`
}

// NewRealm initialize a default realm
func NewRealm(label string) *Realm {
	return &Realm{
		ID:           helpers.IDGeneratorFunc(),
		Label:        label,
		CreationDate: helpers.TimeFunc(),
	}
}

// -----------------------------------------------------------------------------

// Validate entity constraints
func (r *Realm) Validate() error {
	return validation.ValidateStruct(r,
		validation.Field(&r.ID, helpers.IDValidationRules...),
		validation.Field(&r.Label, validation.Required, validation.Length(3, 50), is.PrintableASCII),
	)
}
