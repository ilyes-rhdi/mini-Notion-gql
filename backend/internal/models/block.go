package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type BlockType string

const (
	BlockParagraph BlockType = "PARAGRAPH"
	BlockH1        BlockType = "HEADING1"
	BlockH2        BlockType = "HEADING2"
	BlockTodo      BlockType = "TODO"
	BlockBullet    BlockType = "BULLET"
	BlockNumber    BlockType = "NUMBER"
	BlockQuote     BlockType = "QUOTE"
	BlockCode      BlockType = "CODE"
	BlockImage     BlockType = "IMAGE"
	BlockDivider   BlockType = "DIVIDER"
)

type Block struct {
	ID            string   `gorm:"type:uuid;primaryKey"`
	PageID        string   `gorm:"type:uuid;not null;index"`
	ParentBlockID *string  `gorm:"type:uuid;index"`
	Type          BlockType `gorm:"type:varchar(30);not null"`
	Order         int      `gorm:"not null;default:0"`

	Data datatypes.JSON `gorm:"type:jsonb;not null;default:'{}'"`

	CreatedAt time.Time
	UpdatedAt time.Time
}


func (b *Block) BeforeCreate(tx *gorm.DB) error {
	if b.ID == "" {
		b.ID = uuid.NewString()
	}
	return nil
}
