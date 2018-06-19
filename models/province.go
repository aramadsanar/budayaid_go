package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/uuid"
	"github.com/gobuffalo/validate"
)

type Province struct {
	Id           int    `db:"id" json:"id"`
	Name         string `db:"name" json:"name"`
	FriendlyName string `db:"friendly_name" json:"firendly_name"`
	Island       string `db:"island" json:"island"`
}

// String is not required by pop and may be deleted
func (p Province) String() string {
	jp, _ := json.Marshal(p)
	return string(jp)
}

// Provinces is not required by pop and may be deleted
type Provinces []Province

// String is not required by pop and may be deleted
func (p Provinces) String() string {
	jp, _ := json.Marshal(p)
	return string(jp)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (p *Province) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (p *Province) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (p *Province) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

func (p Province) TableName() string {
	return "province"
}
