/**
Base model
----
Define a minimal commons fields for all models.
- ID , type: uuid.UUID
- CreatedAt , type: time.Time
- UpdatedAt , type: time.Time
- DeletedAt , type: time.Time
Example:
	type Persons struct{
		Base
		Name string
		Age int
	}
*/
package models

import (
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
	"time"
)

// Base contains common columns for all tables.
type Base struct {
	ID        uuid.UUID  `gorm:"type:uuid;primary_key;" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `sql:"index" json:"deleted_at"`
}

// BeforeCreate will set a UUID rather than numeric ID.
func (base *Base) BeforeCreate(scope *gorm.Scope) error {
	_uuid := uuid.NewV4()
	return scope.SetColumn("ID", _uuid)
}