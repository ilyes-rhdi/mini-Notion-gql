package services

import (
	"context"
	"errors"

	"github.com/ilyes-rhdi/buildit-Gql/internal/models"
	"gorm.io/gorm"
)

type PageService struct {
	ws *WorkspaceService
}

func NewPageService() *PageService {
	return &PageService{ws: NewWorkspaceService()}
}

// CreatePage creates a page in a workspace. Parent pages are not exposed in GraphQL v1
// (can be added later by extending the schema).
func (s *PageService) CreatePage(workspaceID, requesterID, title string) (*models.Page, error) {
	ctx := context.Background()

	ok, err := s.ws.isMember(ctx, workspaceID, requesterID)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, errors.New("not authorized")
	}

	if title == "" {
		title = "Untitled"
	}

	p := &models.Page{
		WorkspaceID: workspaceID,
		Title:       title,
		CreatedByID: requesterID,
		Archived:    false,
	}

	if err := getDB().WithContext(ctx).Create(p).Error; err != nil {
		return nil, err
	}
	return p, nil
}

func (s *PageService) GetPage(pageID, requesterID string) (*models.Page, error) {
	ctx := context.Background()

	var p models.Page
	if err := getDB().WithContext(ctx).First(&p, "id = ?", pageID).Error; err != nil {
		return nil, err
	}

	ok, err := s.ws.isMember(ctx, p.WorkspaceID, requesterID)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, errors.New("not authorized")
	}

	return &p, nil
}

func (s *PageService) ListPages(workspaceID, requesterID string) ([]models.Page, error) {
	ctx := context.Background()

	ok, err := s.ws.isMember(ctx, workspaceID, requesterID)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, errors.New("not authorized")
	}

	var pages []models.Page
	if err := getDB().WithContext(ctx).
		Where("workspace_id = ?", workspaceID).
		Order(`"created_at" desc`).
		Find(&pages).Error; err != nil {
		return nil, err
	}
	return pages, nil
}

func (s *PageService) UpdatePage(pageID, requesterID, title string) (*models.Page, error) {
	ctx := context.Background()

	p, err := s.GetPage(pageID, requesterID)
	if err != nil {
		return nil, err
	}

	if title == "" {
		return p, nil
	}

	if err := getDB().WithContext(ctx).
		Model(&models.Page{}).
		Where("id = ?", p.ID).
		Update("title", title).Error; err != nil {
		return nil, err
	}

	if err := getDB().WithContext(ctx).First(p, "id = ?", p.ID).Error; err != nil {
		return nil, err
	}
	return p, nil
}

func (s *PageService) ArchivePage(pageID, requesterID string, archived bool) (*models.Page, error) {
	ctx := context.Background()

	p, err := s.GetPage(pageID, requesterID)
	if err != nil {
		return nil, err
	}

	if err := getDB().WithContext(ctx).
		Model(&models.Page{}).
		Where("id = ?", p.ID).
		Update("archived", archived).Error; err != nil {
		return nil, err
	}

	if err := getDB().WithContext(ctx).First(p, "id = ?", p.ID).Error; err != nil {
		return nil, err
	}
	return p, nil
}

func (s *PageService) DeletePageHard(pageID, requesterID string) error {
	ctx := context.Background()

	p, err := s.GetPage(pageID, requesterID)
	if err != nil {
		return err
	}

	// delete blocks of the page first
	if err := getDB().WithContext(ctx).Where("page_id = ?", p.ID).Delete(&models.Block{}).Error; err != nil {
		return err
	}

	res := getDB().WithContext(ctx).Delete(&models.Page{}, "id = ?", p.ID)
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
