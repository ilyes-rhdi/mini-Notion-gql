package resolvers

import (
	"github.com/graphql-go/graphql"
	"github.com/ilyes-rhdi/buildit-Gql/internal/models"
	"github.com/ilyes-rhdi/buildit-Gql/internal/services"
)

type BlockResolver struct {
	srv *services.BlockService
}

func NewBlockResolver() *BlockResolver {
	return &BlockResolver{srv: services.NewBlockService()}
}

func (r *BlockResolver) Block(p graphql.ResolveParams) (any, error) {
	uid, err := getUserID(p)
	if err != nil {
		return nil, err
	}
	id, _ := p.Args["id"].(string)
	b, err := r.srv.GetBlock(id, uid)
	if err != nil {
		return nil, err
	}
	parent := any(nil)
	if b.ParentBlockID != nil {
		parent = *b.ParentBlockID
	}
	return map[string]any{
		"id":            b.ID,
		"pageId":        b.PageID,
		"parentblockId": parent,
		"Type":          string(b.Type),
		"data":          b.Data,
	}, nil
}

func (r *BlockResolver) Blocks(p graphql.ResolveParams) (any, error) {
	uid, err := getUserID(p)
	if err != nil {
		return nil, err
	}
	pageID, _ := p.Args["pageId"].(string)
	var parentID *string
	if v, ok := p.Args["parentblockId"].(string); ok && v != "" {
		parentID = &v
	}
	items, err := r.srv.ListBlocks(pageID, uid, parentID)
	if err != nil {
		return nil, err
	}
	out := make([]any, 0, len(items))
	for _, b := range items {
		parent := any(nil)
		if b.ParentBlockID != nil {
			parent = *b.ParentBlockID
		}
		out = append(out, map[string]any{
			"id":            b.ID,
			"pageId":        b.PageID,
			"parentblockId": parent,
			"Type":          string(b.Type),
			"data":          b.Data,
		})
	}
	return out, nil
}

func (r *BlockResolver) CreateBlock(p graphql.ResolveParams) (any, error) {
	uid, err := getUserID(p)
	if err != nil {
		return nil, err
	}
	pageID, _ := p.Args["pageId"].(string)
	var parentID *string
	if v, ok := p.Args["parentblockId"].(string); ok && v != "" {
		parentID = &v
	}
	typeStr, _ := p.Args["type"].(string)
	order, hasOrder := p.Args["order"].(int)
	var orderPtr *int
	if hasOrder {
		orderPtr = &order
	}
	data := p.Args["data"]

	b, err := r.srv.CreateBlock(pageID, uid, parentID, models.BlockType(typeStr), orderPtr, data)
	if err != nil {
		return nil, err
	}
	parent := any(nil)
	if b.ParentBlockID != nil {
		parent = *b.ParentBlockID
	}
	return map[string]any{
		"id":            b.ID,
		"pageId":        b.PageID,
		"parentblockId": parent,
		"Type":          string(b.Type),
		"data":          b.Data,
	}, nil
}

func (r *BlockResolver) UpdateBlock(p graphql.ResolveParams) (any, error) {
	uid, err := getUserID(p)
	if err != nil {
		return nil, err
	}
	id, _ := p.Args["id"].(string)
	typeStr, _ := p.Args["type"].(string)
	data := p.Args["data"]

	b, err := r.srv.UpdateBlock(id, uid, models.BlockType(typeStr), data)
	if err != nil {
		return nil, err
	}
	parent := any(nil)
	if b.ParentBlockID != nil {
		parent = *b.ParentBlockID
	}
	return map[string]any{
		"id":            b.ID,
		"pageId":        b.PageID,
		"parentblockId": parent,
		"Type":          string(b.Type),
		"data":          b.Data,
	}, nil
}

func (r *BlockResolver) MoveBlock(p graphql.ResolveParams) (any, error) {
	uid, err := getUserID(p)
	if err != nil {
		return nil, err
	}
	id, _ := p.Args["id"].(string)
	pageID, _ := p.Args["pageId"].(string)
	var parentID *string
	if v, ok := p.Args["parentblockId"].(string); ok && v != "" {
		parentID = &v
	}
	order, _ := p.Args["order"].(int)

	b, err := r.srv.MoveBlock(id, uid, pageID, parentID, order)
	if err != nil {
		return nil, err
	}
	parent := any(nil)
	if b.ParentBlockID != nil {
		parent = *b.ParentBlockID
	}
	return map[string]any{
		"id":            b.ID,
		"pageId":        b.PageID,
		"parentblockId": parent,
		"Type":          string(b.Type),
		"data":          b.Data,
	}, nil
}

func (r *BlockResolver) DeleteBlockTree(p graphql.ResolveParams) (any, error) {
	uid, err := getUserID(p)
	if err != nil {
		return nil, err
	}
	id, _ := p.Args["id"].(string)
	if err := r.srv.DeleteBlockTree(id, uid); err != nil {
		return nil, err
	}
	return true, nil
}
