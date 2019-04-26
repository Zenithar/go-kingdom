package models_test

import (
	"testing"

	. "github.com/onsi/gomega"

	"go.zenithar.org/kingdom/internal/models"
)

func TestRealmCreation(t *testing.T) {
	g := NewGomegaWithT(t)

	entity := models.NewRealm("foo")
	g.Expect(entity).ToNot(BeNil(), "Entity should not be nil")
	g.Expect(entity.Validate()).To(BeNil(), "Validation error should be nil")
	g.Expect(entity.Label).ToNot(BeEmpty(), "Label should not be blank")
	g.Expect(entity.Label).To(Equal("foo"), "Label should be as expected")
}

func TestRealmValidation(t *testing.T) {
	g := NewGomegaWithT(t)

	for _, f := range []struct {
		label     string
		expectErr bool
	}{
		{"a", true},
		{"aa", true},
		{"foo", false},
		{"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", false},
		{"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", true},
	} {
		obj := models.NewRealm(f.label)
		g.Expect(obj).ToNot(BeNil(), "Entity should not be nil")

		if err := obj.Validate(); err != nil {
			if !f.expectErr {
				t.Errorf("Validation error should not be raised, %v raised", err)
			}
		} else {
			if f.expectErr {
				t.Error("Validation error should be raised")
			}
		}
	}
}
