package models

import (
	"time"
)

// Role represents the database model of roles
type Role struct {
	ID          uint   `gorm:"primary_key" json:"id"`
	Name        string `gorm:"size:255;not null" json:"name"`
	GuardName   string `gorm:"size:255;not null;uniqueIndex" json:"guard_name"`
	Description string `gorm:"size:255;" json:"description"`

	// Many to Many
	Permissions []Permission `gorm:"many2many:role_permissions;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"permissions"`

	// Time
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// TableName sets the table name
func (Role) TableName() string {
	return "roles"
}
