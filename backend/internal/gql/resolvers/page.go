package resolvers

import (
	"github.com/graphql-go/graphql"
	"github.com/ilyes-rhdi/buildit-Gql/internal/services"
)

type PageResolver struct {
	srv *services.PageService
}

func NewPageResolver() *PageResolver {
	return &PageResolver{srv: services.NewPageService()}
}

func (r *PageResolver) Page(p graphql.ResolveParams) (any, error) {
	uid, err := getUserID(p)
	if err != nil {
		return nil, err
	}
	id, _ := p.Args["id"].(string)
	pg, err := r.srv.GetPage(id, uid)
	if err != nil {
		return nil, err
	}
	return pageToMap(*pg), nil
}

func (r *PageResolver) Pages(p graphql.ResolveParams) (any, error) {
	uid, err := getUserID(p)
	if err != nil {
		return nil, err
	}
	workspaceID, _ := p.Args["workspaceId"].(string)
	items, err := r.srv.ListPages(workspaceID, uid)
	if err != nil {
		return nil, err
	}
	out := make([]any, 0, len(items))
	for _, pg := range items {
		out = append(out, pageToMap(pg))
	}
	return out, nil
}

func (r *PageResolver) CreatePage(p graphql.ResolveParams) (any, error) {
	uid, err := getUserID(p)
	if err != nil {
		return nil, err
	}
	workspaceID, _ := p.Args["workspaceId"].(string)
	title, _ := p.Args["title"].(string)
	pg, err := r.srv.CreatePage(workspaceID, uid, title)
	if err != nil {
		return nil, err
	}
	return pageToMap(*pg), nil
}

func (r *PageResolver) UpdatePage(p graphql.ResolveParams) (any, error) {
	uid, err := getUserID(p)
	if err != nil {
		return nil, err
	}
	id, _ := p.Args["id"].(string)
	title, _ := p.Args["title"].(string)
	pg, err := r.srv.UpdatePage(id, uid, title)
	if err != nil {
		return nil, err
	}
	return pageToMap(*pg), nil
}

func (r *PageResolver) ArchivePage(p graphql.ResolveParams) (any, error) {
	uid, err := getUserID(p)
	if err != nil {
		return nil, err
	}
	id, _ := p.Args["id"].(string)
	archived, _ := p.Args["archived"].(bool)
	pg, err := r.srv.ArchivePage(id, uid, archived)
	if err != nil {
		return nil, err
	}
	return pageToMap(*pg), nil
}

func (r *PageResolver) DeletePageHard(p graphql.ResolveParams) (any, error) {
	uid, err := getUserID(p)
	if err != nil {
		return nil, err
	}
	id, _ := p.Args["id"].(string)
	if err := r.srv.DeletePageHard(id, uid); err != nil {
		return nil, err
	}
	return true, nil
}
