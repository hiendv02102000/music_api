package entity

import (
	"time"
)

type roleAllowed string

const (
	AdminRole  roleAllowed = "admin"
	ClientRole roleAllowed = "client"
)

// Users struct
type Users struct {
	ID             int         `gorm:"column:id;primary_key;auto_increment;not null"`
	FirstName      string      `gorm:"column:first_name;"`
	LastName       string      `gorm:"column:last_name"`
	Username       string      `gorm:"column:username;not null"`
	Password       string      `gorm:"column:password;not null"`
	Role           roleAllowed `gorm:"column:role"`
	Token          *string     `gorm:"column:token"`
	TokenExpriedAt *time.Time  `gorm:"column:token_expired_at"`
	Novels         []Novels
	BaseModel
}
