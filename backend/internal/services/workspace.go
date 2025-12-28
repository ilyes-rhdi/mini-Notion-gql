package services

import (
	"context"
	"errors"

	"github.com/ilyes-rhdi/buildit-Gql/internal/models"
	"gorm.io/gorm"
)



type WorkspaceService struct{}

func NewWorkspaceService() *WorkspaceService {
	return &WorkspaceService{}
}

func (s *WorkspaceService) CreateWorkspace(name, ownerID string) (*models.Workspace, error) {
	ctx := context.Background()
	if name == "" {
		name = "My Workspace"
	}

	ws := &models.Workspace{
		Name:    name,
		OwnerID: ownerID,
	}

	err := getDB().WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(ws).Error; err != nil {
			return err
		}

		// ajoute le owner dans members
		member := &models.WorkspaceMember{
			WorkspaceID: ws.ID,
			UserID:      ownerID,
			Role:        models.RoleOwner,
		}
		if err := tx.Create(member).Error; err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	// preload utile pour GraphQL
	if err := getDB().WithContext(ctx).
		Preload("Owner").
		Preload("Members").
		Preload("Members.User").
		First(ws, "id = ?", ws.ID).Error; err != nil {
		return nil, err
	}

	return ws, nil
}

func (s *WorkspaceService) GetWorkspace(workspaceID, requesterID string) (*models.Workspace, error) {
	ctx := context.Background()

	ok, err := s.isMember(ctx, workspaceID, requesterID)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, errors.New("not authorized")
	}

	var ws models.Workspace
	if err := getDB().WithContext(ctx).
		Preload("Owner").
		Preload("Members").
		Preload("Members.User").
		Preload("Pages").
		First(&ws, "id = ?", workspaceID).Error; err != nil {
		return nil, err
	}
	return &ws, nil
}

func (s *WorkspaceService) ListWorkspaces(requesterID string) ([]models.Workspace, error) {
	ctx := context.Background()

	var memberships []models.WorkspaceMember
	if err := getDB().WithContext(ctx).
		Where("user_id = ?", requesterID).
		Find(&memberships).Error; err != nil {
		return nil, err
	}

	ids := make([]string, 0, len(memberships))
	for _, m := range memberships {
		ids = append(ids, m.WorkspaceID)
	}
	if len(ids) == 0 {
		return []models.Workspace{}, nil
	}

	var workspaces []models.Workspace
	if err := getDB().WithContext(ctx).
		Where("id IN ?", ids).
		Preload("Owner").
		Find(&workspaces).Error; err != nil {
		return nil, err
	}
	return workspaces, nil
}

func (s *WorkspaceService) AddMember(workspaceID, requesterID, userID string, role models.WorkspaceRole) (*models.WorkspaceMember, error) {
	ctx := context.Background()

	// politique simple: seul owner peut ajouter
	if err := s.requireOwner(ctx, workspaceID, requesterID); err != nil {
		return nil, err
	}

	// user doit exister
	var u models.User
	if err := getDB().WithContext(ctx).First(&u, "id = ?", userID).Error; err != nil {
		return nil, err
	}

	if role == "" {
		role = models.RoleMember
	}
	// évite d’ajouter OWNER via cette action
	if role != models.RoleAdmin && role != models.RoleMember {
		role = models.RoleMember
	}

	// pas de doublon
	var existing models.WorkspaceMember
	err := getDB().WithContext(ctx).
		Where("workspace_id = ? AND user_id = ?", workspaceID, userID).
		First(&existing).Error
	if err == nil {
		return nil, errors.New("user already member")
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	m := &models.WorkspaceMember{
		WorkspaceID: workspaceID,
		UserID:      userID,
		Role:        role,
	}
	if err := getDB().WithContext(ctx).Create(m).Error; err != nil {
		return nil, err
	}

	// preload user pour GraphQL
	_ = getDB().WithContext(ctx).Preload("User").First(m, "id = ?", m.ID).Error
	return m, nil
}

func (s *WorkspaceService) UpdateMemberRole(workspaceID, requesterID, targetUserID string, role models.WorkspaceRole) (*models.WorkspaceMember, error) {
	ctx := context.Background()

	// politique simple: seul owner peut changer rôles
	if err := s.requireOwner(ctx, workspaceID, requesterID); err != nil {
		return nil, err
	}

	// interdire ROLE OWNER ici
	if role != models.RoleAdmin && role != models.RoleMember {
		return nil, errors.New("invalid role")
	}

	// owner ne change pas via cette mutation
	var ws models.Workspace
	if err := getDB().WithContext(ctx).Select("id", "owner_id").First(&ws, "id = ?", workspaceID).Error; err != nil {
		return nil, err
	}
	if targetUserID == ws.OwnerID {
		return nil, errors.New("cannot change owner role here")
	}

	var m models.WorkspaceMember
	if err := getDB().WithContext(ctx).
		Where("workspace_id = ? AND user_id = ?", workspaceID, targetUserID).
		First(&m).Error; err != nil {
		return nil, err
	}

	if err := getDB().WithContext(ctx).Model(&m).Update("role", role).Error; err != nil {
		return nil, err
	}
	_ = getDB().WithContext(ctx).Preload("User").First(&m, "id = ?", m.ID).Error
	return &m, nil
}

func (s *WorkspaceService) RemoveMember(workspaceID, requesterID, targetUserID string) error {
	ctx := context.Background()

	// user peut quitter lui-même
	if requesterID != targetUserID {
		// sinon seul owner peut remove quelqu’un
		if err := s.requireOwner(ctx, workspaceID, requesterID); err != nil {
			return err
		}
	}

	// ne pas remove le owner via removeMember (utilise transferOwnership)
	var ws models.Workspace
	if err := getDB().WithContext(ctx).Select("id", "owner_id").First(&ws, "id = ?", workspaceID).Error; err != nil {
		return err
	}
	if targetUserID == ws.OwnerID {
		return errors.New("cannot remove owner; transfer ownership first")
	}

	return getDB().WithContext(ctx).
		Where("workspace_id = ? AND user_id = ?", workspaceID, targetUserID).
		Delete(&models.WorkspaceMember{}).Error
}

func (s *WorkspaceService) TransferOwnership(workspaceID, requesterID, newOwnerUserID string) (*models.Workspace, error) {
	ctx := context.Background()

	// seul owner
	if err := s.requireOwner(ctx, workspaceID, requesterID); err != nil {
		return nil, err
	}

	err := getDB().WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		var ws models.Workspace
		if err := tx.First(&ws, "id = ?", workspaceID).Error; err != nil {
			return err
		}

		// new owner doit être membre (sinon ajoute-le)
		var newMem models.WorkspaceMember
		err := tx.Where("workspace_id = ? AND user_id = ?", workspaceID, newOwnerUserID).First(&newMem).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			newMem = models.WorkspaceMember{WorkspaceID: workspaceID, UserID: newOwnerUserID, Role: models.RoleAdmin}
			if err := tx.Create(&newMem).Error; err != nil {
				return err
			}
		} else if err != nil {
			return err
		}

		// downgrade ancien owner en ADMIN
		if err := tx.Model(&models.WorkspaceMember{}).
			Where("workspace_id = ? AND user_id = ?", workspaceID, ws.OwnerID).
			Update("role", models.RoleAdmin).Error; err != nil {
			return err
		}

		// set new owner
		if err := tx.Model(&models.Workspace{}).
			Where("id = ?", workspaceID).
			Update("owner_id", newOwnerUserID).Error; err != nil {
			return err
		}

		// upgrade role membership new owner
		if err := tx.Model(&models.WorkspaceMember{}).
			Where("workspace_id = ? AND user_id = ?", workspaceID, newOwnerUserID).
			Update("role", models.RoleOwner).Error; err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	// retourne workspace à jour
	var ws models.Workspace
	if err := getDB().WithContext(ctx).
		Preload("Owner").
		Preload("Members").
		Preload("Members.User").
		Preload("Pages").
		First(&ws, "id = ?", workspaceID).Error; err != nil {
		return nil, err
	}
	return &ws, nil
}

func (s *WorkspaceService) isMember(ctx context.Context, workspaceID, userID string) (bool, error) {
	var m models.WorkspaceMember
	err := getDB().WithContext(ctx).
		Where("workspace_id = ? AND user_id = ?", workspaceID, userID).
		First(&m).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false, nil
	}
	return err == nil, err
}

func (s *WorkspaceService) requireOwner(ctx context.Context, workspaceID, userID string) error {
	var ws models.Workspace
	if err := getDB().WithContext(ctx).Select("id", "owner_id").First(&ws, "id = ?", workspaceID).Error; err != nil {
		return err
	}
	if ws.OwnerID != userID {
		return errors.New("not authorized")
	}
	return nil
}
