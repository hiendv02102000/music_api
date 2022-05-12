package entity

import (
	"time"

	"github.com/jinzhu/gorm"
)

// BaseModel struct
type BaseModel struct {
	CreatedAt time.Time  `gorm:"column:created_at"`
	UpdatedAt time.Time  `gorm:"column:updated_at"`
	DeletedAt *time.Time `gorm:"column:deleted_at"`
}

func (m *BaseModel) BeforeCreate(db *gorm.DB) error {
	m.CreatedAt = time.Now().Local()
	return nil
}

func (m *BaseModel) BeforeUpdate(db *gorm.DB) error {
	m.UpdatedAt = time.Now().Local()
	return nil
}
func (m *BaseModel) BeforeDelete(db *gorm.DB) error {
	m.UpdatedAt = time.Now().Local()
	return nil
}
