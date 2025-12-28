package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID        string `gorm:"type:uuid;primaryKey"`
	Email     string `gorm:"uniqueIndex;not null"`
	Name      string `gorm:"not null"`
	Password  string `gorm:"not null"`
	Active    bool   `gorm:"default:false"`

	Image  string `gorm:"not null;default:'uploads/profiles/default.jpg'"`
	Bio    string `gorm:"not null;default:''"`
	Adress        string `gorm:"not null;default:''"`
	Phone         string `gorm:"not null;default:''"`
	ExternalLinks string `gorm:"not null;default:''"`
	Gender *bool   `gorm:"not null;default:false"`
	BgImg  string `gorm:"not null;default:'uploads/bgs/default.jpg'"`

	CreatedAt time.Time
	UpdatedAt time.Time
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	if u.ID == "" {
		u.ID = uuid.NewString()
	}
	return nil
}
