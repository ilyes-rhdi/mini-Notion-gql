package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Page struct {
	ID          string `gorm:"type:uuid;primaryKey"`
	WorkspaceID string `gorm:"type:uuid;not null;index"`

	Title    string `gorm:"not null;default:'Untitled'"`
	Icon     string
	Cover    string
	Archived bool `gorm:"not null;default:false"`

	ParentPageID *string `gorm:"type:uuid;index"`
	CreatedByID  string  `gorm:"type:uuid;not null;index"`

	Workspace Workspace `gorm:"foreignKey:WorkspaceID"`
	CreatedBy User      `gorm:"foreignKey:CreatedByID"`
	Blocks    []Block   `gorm:"foreignKey:PageID"`

	CreatedAt time.Time
	UpdatedAt time.Time
}

func (p *Page) BeforeCreate(tx *gorm.DB) error {
	if p.ID == "" {
		p.ID = uuid.NewString()
	}
	return nil
}
