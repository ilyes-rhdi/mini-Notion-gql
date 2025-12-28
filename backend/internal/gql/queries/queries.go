package queries

import (
	"github.com/graphql-go/graphql"
	"github.com/ilyes-rhdi/buildit-Gql/internal/gql/resolvers"
	"github.com/ilyes-rhdi/buildit-Gql/internal/gql/types"
)

var RootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name: "Query",
	Fields: graphql.Fields{
		"me": &graphql.Field{
			Type:    types.UserType,
			Resolve: resolvers.User.Me,
		},
		"workspace": &graphql.Field{
			Type: types.WorkspaceType,
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
			},
			Resolve: resolvers.Workspace.Workspace,
		},
		"workspaces": &graphql.Field{
			Type:    graphql.NewList(types.WorkspaceType),
			Resolve: resolvers.Workspace.Workspaces,
		},
		"page": &graphql.Field{
			Type: types.PageType,
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
			},
			Resolve: resolvers.Page.Page,
		},
		"pages": &graphql.Field{
			Type: graphql.NewList(types.PageType),
			Args: graphql.FieldConfigArgument{
				"workspaceId": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
			},
			Resolve: resolvers.Page.Pages,
		},
		"block": &graphql.Field{
			Type: types.BlockType,
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
			},
			Resolve: resolvers.Block.Block,
		},
		"blocks": &graphql.Field{
			Type: graphql.NewList(types.BlockType),
			Args: graphql.FieldConfigArgument{
				"pageId":        &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
				"parentblockId": &graphql.ArgumentConfig{Type: graphql.String},
			},
			Resolve: resolvers.Block.Blocks,
		},
	},
})
