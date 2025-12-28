package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Workspace struct {
	ID      string `gorm:"type:uuid;primaryKey"`
	Name    string `gorm:"not null"`
	OwnerID string `gorm:"type:uuid;not null;index"`
	Owner   User              `gorm:"foreignKey:OwnerID"`
	Members []WorkspaceMember `gorm:"foreignKey:WorkspaceID"`
	Pages   []Page            `gorm:"foreignKey:WorkspaceID"`

	CreatedAt time.Time
	UpdatedAt time.Time
}

func (w *Workspace) BeforeCreate(tx *gorm.DB) error {
	if w.ID == "" {
		w.ID = uuid.NewString()
	}
	return nil
}

type WorkspaceRole string

const (
	RoleOwner  WorkspaceRole = "OWNER"
	RoleAdmin  WorkspaceRole = "ADMIN"
	RoleMember WorkspaceRole = "MEMBER"
)

type WorkspaceMember struct {
	ID          string        `gorm:"type:uuid;primaryKey"`
	WorkspaceID string        `gorm:"type:uuid;not null;index"`
	UserID      string        `gorm:"type:uuid;not null;index"`
	Role        WorkspaceRole `gorm:"type:varchar(20);not null;default:'MEMBER'"`

	Workspace Workspace `gorm:"foreignKey:WorkspaceID"`
	User      User      `gorm:"foreignKey:UserID"`

	CreatedAt time.Time
	UpdatedAt time.Time
}

func (m *WorkspaceMember) BeforeCreate(tx *gorm.DB) error {
	if m.ID == "" {
		m.ID = uuid.NewString()
	}
	return nil
}
