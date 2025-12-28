package resolvers

import (
	"github.com/graphql-go/graphql"
	"github.com/ilyes-rhdi/buildit-Gql/internal/models"
	"github.com/ilyes-rhdi/buildit-Gql/internal/services"
)

type WorkspaceResolver struct {
	srv *services.WorkspaceService
}

func NewWorkspaceResolver() *WorkspaceResolver {
	return &WorkspaceResolver{srv: services.NewWorkspaceService()}
}

func (r *WorkspaceResolver) Workspace(p graphql.ResolveParams) (any, error) {
	uid, err := getUserID(p)
	if err != nil {
		return nil, err
	}
	id, _ := p.Args["id"].(string)
	ws, err := r.srv.GetWorkspace(id, uid)
	if err != nil {
		return nil, err
	}
	return workspaceToMap(*ws), nil
}

func (r *WorkspaceResolver) Workspaces(p graphql.ResolveParams) (any, error) {
	uid, err := getUserID(p)
	if err != nil {
		return nil, err
	}
	items, err := r.srv.ListWorkspaces(uid)
	if err != nil {
		return nil, err
	}
	out := make([]any, 0, len(items))
	for _, ws := range items {
		out = append(out, workspaceToMap(ws))
	}
	return out, nil
}

func (r *WorkspaceResolver) CreateWorkspace(p graphql.ResolveParams) (any, error) {
	uid, err := getUserID(p)
	if err != nil {
		return nil, err
	}
	name, _ := p.Args["name"].(string)
	ws, err := r.srv.CreateWorkspace(name, uid)
	if err != nil {
		return nil, err
	}
	return workspaceToMap(*ws), nil
}

func (r *WorkspaceResolver) AddMember(p graphql.ResolveParams) (any, error) {
	uid, err := getUserID(p)
	if err != nil {
		return nil, err
	}
	workspaceID, _ := p.Args["workspaceId"].(string)
	userID, _ := p.Args["userId"].(string)
	roleStr, _ := p.Args["role"].(string)
	role := models.WorkspaceRole(roleStr)

	m, err := r.srv.AddMember(workspaceID, uid, userID, role)
	if err != nil {
		return nil, err
	}
	return memberToMap(*m), nil
}

func (r *WorkspaceResolver) UpdateMemberRole(p graphql.ResolveParams) (any, error) {
	uid, err := getUserID(p)
	if err != nil {
		return nil, err
	}
	workspaceID, _ := p.Args["workspaceId"].(string)
	userID, _ := p.Args["userId"].(string)
	roleStr, _ := p.Args["role"].(string)
	role := models.WorkspaceRole(roleStr)

	m, err := r.srv.UpdateMemberRole(workspaceID, uid, userID, role)
	if err != nil {
		return nil, err
	}
	return memberToMap(*m), nil
}

func (r *WorkspaceResolver) RemoveMember(p graphql.ResolveParams) (any, error) {
	uid, err := getUserID(p)
	if err != nil {
		return nil, err
	}
	workspaceID, _ := p.Args["workspaceId"].(string)
	userID, _ := p.Args["userId"].(string)

	if err := r.srv.RemoveMember(workspaceID, uid, userID); err != nil {
		return nil, err
	}
	return true, nil
}

func (r *WorkspaceResolver) TransferOwnership(p graphql.ResolveParams) (any, error) {
	uid, err := getUserID(p)
	if err != nil {
		return nil, err
	}
	workspaceID, _ := p.Args["workspaceId"].(string)
	newOwnerUserID, _ := p.Args["newOwnerUserId"].(string)

	ws, err := r.srv.TransferOwnership(workspaceID, uid, newOwnerUserID)
	if err != nil {
		return nil, err
	}
	return workspaceToMap(*ws), nil
}
