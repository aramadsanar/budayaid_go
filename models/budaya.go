package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/uuid"
	"github.com/gobuffalo/validate"
)

type Budaya struct {
	Id               int    `db:"id" json:"id"`
	Name             string `db:"name" json:"name"`
	Description      string `db:"description" json:"description"`
	ImageUrl         string `db:"image_url" json:"image_url"`
	GoogleSearchTerm string `db:"google_search_term" json:"google_search_term"`

	Province   Province   `has_one:"provinces" fk_id:"id"`
	Categories Categories `has_one:"categories" fk_id:"id"`
}

// String is not required by pop and may be deleted
func (b Budaya) String() string {
	jb, _ := json.Marshal(b)
	return string(jb)
}

// Budayas is not required by pop and may be deleted
type Budayas []Budaya

// String is not required by pop and may be deleted
func (b Budayas) String() string {
	jb, _ := json.Marshal(b)
	return string(jb)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (b *Budaya) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (b *Budaya) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (b *Budaya) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

func (b Budaya) TableName() string {
	return "budaya"
}
