package resolvers

import (
	"errors"

	"github.com/graphql-go/graphql"
	middlewares "github.com/ilyes-rhdi/buildit-Gql/internal/middlewares/gql"
	"github.com/ilyes-rhdi/buildit-Gql/internal/models"
)

func userToMap(u models.User) map[string]any {
	return map[string]any{
		"id":     u.ID,
		"email":  u.Email,
		"name":   u.Name,
		"bio":    u.Bio,
		"joined": u.CreatedAt,
		"image":  u.Image,
		"gender": u.Gender,
		// champs optionnels (même si pas dans UserType)
		"adress": u.Adress,
		"phone":  u.Phone,
		"links":  u.ExternalLinks,
	}
}

func pageToMap(p models.Page) map[string]any {
	return map[string]any{
		"id":          p.ID,
		"workspaceId": p.WorkspaceID,
		"title":       p.Title,
		"archived":    p.Archived,
	}
}

func memberToMap(m models.WorkspaceMember) map[string]any {
	out := map[string]any{
		"id":          m.ID,
		"workspaceId": m.WorkspaceID,
		"userId":      m.UserID,
		"role":        string(m.Role),
	}
	// si preload User
	if m.User.ID != "" {
		out["user"] = userToMap(m.User)
	}
	return out
}

func workspaceToMap(ws models.Workspace) map[string]any {
	out := map[string]any{
		"id":      ws.ID,
		"name":    ws.Name,
		"OwnerID": ws.OwnerID,
	}
	if ws.Owner.ID != "" {
		out["Owner"] = userToMap(ws.Owner)
	}
	if ws.Members != nil {
		members := make([]any, 0, len(ws.Members))
		for _, m := range ws.Members {
			members = append(members, memberToMap(m))
		}
		out["Members"] = members
	}
	if ws.Pages != nil {
		pages := make([]any, 0, len(ws.Pages))
		for _, p := range ws.Pages {
			pages = append(pages, pageToMap(p))
		}
		out["Pages"] = pages
	}
	return out
}

func getUserID(p graphql.ResolveParams) (string, error) {
	uid, err := middlewares.IsAuthenticated(p)
	if err != nil {
		return "", err
	}
	if uid == "" {
		return "", errors.New("Not Authorized")
	}
	return uid, nil
}
