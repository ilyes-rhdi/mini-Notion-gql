package services

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/ilyes-rhdi/buildit-Gql/internal/models"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type BlockService struct {
	pg *PageService
}

func NewBlockService() *BlockService {
	return &BlockService{
		pg: NewPageService(),
	}
}

func (s *BlockService) CreateBlock(pageID, requesterID string, parentBlockID *string, blockType models.BlockType, order *int, data any) (*models.Block, error) {
	ctx := context.Background()

	// access check via page
	_, err := s.pg.GetPage(pageID, requesterID)
	if err != nil {
		return nil, err
	}

	b := &models.Block{
		PageID:        pageID,
		ParentBlockID: parentBlockID,
		Type:          blockType,
		Order:         0,
		Data:          datatypes.JSON([]byte(`{}`)),
	}

	if order != nil {
		b.Order = *order
	} else {
		// order auto = max(order)+1 among siblings
		var maxOrder int
		q := getDB().WithContext(ctx).Model(&models.Block{}).Where("page_id = ?", pageID)
		if parentBlockID == nil {
			q = q.Where("parent_block_id IS NULL")
		} else {
			q = q.Where("parent_block_id = ?", *parentBlockID)
		}
		_ = q.Select("COALESCE(MAX(\"order\"), 0)").Scan(&maxOrder).Error
		b.Order = maxOrder + 1
	}

	if data != nil {
		buf, err := json.Marshal(data)
		if err != nil {
			return nil, err
		}
		b.Data = datatypes.JSON(buf)
	}

	if err := getDB().WithContext(ctx).Create(b).Error; err != nil {
		return nil, err
	}

	return b, nil
}

func (s *BlockService) GetBlock(blockID, requesterID string) (*models.Block, error) {
	ctx := context.Background()

	var b models.Block
	if err := getDB().WithContext(ctx).First(&b, "id = ?", blockID).Error; err != nil {
		return nil, err
	}

	// access check via page
	_, err := s.pg.GetPage(b.PageID, requesterID)
	if err != nil {
		return nil, err
	}

	return &b, nil
}

func (s *BlockService) ListBlocks(pageID, requesterID string, parentBlockID *string) ([]models.Block, error) {
	ctx := context.Background()

	// access check via page
	_, err := s.pg.GetPage(pageID, requesterID)
	if err != nil {
		return nil, err
	}

	q := getDB().WithContext(ctx).Where("page_id = ?", pageID)
	if parentBlockID == nil {
		q = q.Where("parent_block_id IS NULL")
	} else {
		q = q.Where("parent_block_id = ?", *parentBlockID)
	}

	var blocks []models.Block
	if err := q.Order(`"order" asc`).Find(&blocks).Error; err != nil {
		return nil, err
	}
	return blocks, nil
}

// UpdateBlock updates (type, data). GraphQL requires `type` non-null.
func (s *BlockService) UpdateBlock(blockID, requesterID string, blockType models.BlockType, data any) (*models.Block, error) {
	ctx := context.Background()

	b, err := s.GetBlock(blockID, requesterID)
	if err != nil {
		return nil, err
	}

	updates := map[string]any{}
	if blockType != "" {
		updates["type"] = blockType
	}
	if data != nil {
		buf, err := json.Marshal(data)
		if err != nil {
			return nil, err
		}
		updates["data"] = datatypes.JSON(buf)
	}
	if len(updates) == 0 {
		return b, nil
	}

	if err := getDB().WithContext(ctx).Model(&models.Block{}).
		Where("id = ?", b.ID).
		Updates(updates).Error; err != nil {
		return nil, err
	}

	if err := getDB().WithContext(ctx).First(b, "id = ?", b.ID).Error; err != nil {
		return nil, err
	}
	return b, nil
}

// MoveBlock moves a block to another page / parent and reorders it.
func (s *BlockService) MoveBlock(blockID, requesterID, newPageID string, newParentBlockID *string, newOrder int) (*models.Block, error) {
	ctx := context.Background()

	b, err := s.GetBlock(blockID, requesterID)
	if err != nil {
		return nil, err
	}

	// access check on target page
	_, err = s.pg.GetPage(newPageID, requesterID)
	if err != nil {
		return nil, err
	}

	err = getDB().WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// shift siblings in target list to make room
		q := tx.Model(&models.Block{}).Where("page_id = ?", newPageID)
		if newParentBlockID == nil {
			q = q.Where("parent_block_id IS NULL")
		} else {
			q = q.Where("parent_block_id = ?", *newParentBlockID)
		}
		if err := q.Where(`"order" >= ?`, newOrder).
			Update(`"order"`, gorm.Expr(`"order" + 1`)).Error; err != nil {
			return err
		}

		updates := map[string]any{
			"page_id":         newPageID,
			"parent_block_id": newParentBlockID,
			"order":           newOrder,
		}
		if err := tx.Model(&models.Block{}).Where("id = ?", b.ID).Updates(updates).Error; err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	if err := getDB().WithContext(ctx).First(b, "id = ?", b.ID).Error; err != nil {
		return nil, err
	}
	return b, nil
}

func (s *BlockService) DeleteBlockTree(blockID, requesterID string) error {
	ctx := context.Background()

	b, err := s.GetBlock(blockID, requesterID)
	if err != nil {
		return err
	}

	// BFS recursion
	toVisit := []string{b.ID}
	all := make([]string, 0, 16)
	seen := map[string]bool{}

	for len(toVisit) > 0 {
		cur := toVisit[0]
		toVisit = toVisit[1:]
		if seen[cur] {
			continue
		}
		seen[cur] = true
		all = append(all, cur)

		var children []models.Block
		if err := getDB().WithContext(ctx).
			Select("id").
			Where("parent_block_id = ?", cur).
			Find(&children).Error; err != nil {
			return err
		}
		for _, ch := range children {
			toVisit = append(toVisit, ch.ID)
		}
	}

	if len(all) == 0 {
		return nil
	}

	if err := getDB().WithContext(ctx).Where("id IN ?", all).Delete(&models.Block{}).Error; err != nil {
		return err
	}
	return nil
}

func (s *BlockService) DeleteBlockHard(blockID, requesterID string) error {
	ctx := context.Background()

	_, err := s.GetBlock(blockID, requesterID)
	if err != nil {
		return err
	}

	res := getDB().WithContext(ctx).Delete(&models.Block{}, "id = ?", blockID)
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return errors.New("block not found")
	}
	return nil
}
