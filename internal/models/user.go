package models

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"

	"go.zenithar.org/kingdom/internal/helpers"
)

// User represents an user identity in the realm
type User struct {
	RealmID      string    `json:"realm_id" db:"realm_id"`
	ID           string    `json:"id"  db:"user_id"`
	Principal    string    `json:"principal" db:"principal"`
	Secret       string    `json:"secret" db:"secret"`
	CreationDate time.Time `json:"creation_date" db:"creation_date"`
}

// NewUser returns an initialized user account
func NewUser(realmID string) *User {
	return &User{
		RealmID:      realmID,
		ID:           helpers.IDGeneratorFunc(),
		CreationDate: helpers.TimeFunc(),
	}
}

// -----------------------------------------------------------------------------

// Validate entity constraints
func (u *User) Validate() error {
	return validation.ValidateStruct(u,
		validation.Field(&u.RealmID, helpers.IDValidationRules...),
		validation.Field(&u.ID, helpers.IDValidationRules...),
		validation.Field(&u.Principal, validation.Required, is.PrintableASCII),
		validation.Field(&u.Secret, validation.Required, is.PrintableASCII),
	)
}

// SetPrincipal is used to update principal hash
func (u *User) SetPrincipal(principal string) {
	u.Principal = helpers.PrincipalHashFunc(principal)
}

// Authenticate user identity
func (u *User) Authenticate(password string) (bool, error) {
	return helpers.CheckPasswordFunc(u.Secret, password)
}
