package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

//Base base struct
type Base struct {
	ID        uuid.UUID      `json:"id" gorm:"primary_key"`
	CreatedAt time.Time      `json:"created_at" gorm:"column:created_at"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"column:updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"column:deleted_at"`
}

// BeforeCreate will set a UUID rather than numeric ID.
func (base *Base) BeforeCreate(_ *gorm.DB) error {
	if base.ID == uuid.Nil {
		base.ID = uuid.New()
	}
	if base.CreatedAt.IsZero() {
		base.CreatedAt = time.Now()
	}
	if base.UpdatedAt.IsZero() {
		base.UpdatedAt = time.Now()
	}
	return nil
}
