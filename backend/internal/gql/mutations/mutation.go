package mutations

import (
	"github.com/graphql-go/graphql"
	"github.com/ilyes-rhdi/buildit-Gql/internal/gql/resolvers"
	"github.com/ilyes-rhdi/buildit-Gql/internal/gql/types"
)

var RootMutation = graphql.NewObject(graphql.ObjectConfig{
	Name: "Mutation",
	Fields: graphql.Fields{
		"createWorkspace": &graphql.Field{
			Type: types.WorkspaceType,
			Args: graphql.FieldConfigArgument{
				"name": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
			},
			Resolve: resolvers.Workspace.CreateWorkspace,
		},
		"addWorkspaceMember": &graphql.Field{
			Type: types.WorkspaceMemberType,
			Args: graphql.FieldConfigArgument{
				"workspaceId": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
				"userId":      &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
				"role":        &graphql.ArgumentConfig{Type: types.WorkspaceRoleEnum},
			},
			Resolve: resolvers.Workspace.AddMember,
		},
		"updateWorkspaceMemberRole": &graphql.Field{
			Type: types.WorkspaceMemberType,
			Args: graphql.FieldConfigArgument{
				"workspaceId": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
				"userId":      &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
				"role":        &graphql.ArgumentConfig{Type: graphql.NewNonNull(types.WorkspaceRoleEnum)},
			},
			Resolve: resolvers.Workspace.UpdateMemberRole,
		},
		"removeWorkspaceMember": &graphql.Field{
			Type: graphql.Boolean,
			Args: graphql.FieldConfigArgument{
				"workspaceId": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
				"userId":      &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
			},
			Resolve: resolvers.Workspace.RemoveMember,
		},
		"transferWorkspaceOwnership": &graphql.Field{
			Type: types.WorkspaceType,
			Args: graphql.FieldConfigArgument{
				"workspaceId":    &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
				"newOwnerUserId": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
			},
			Resolve: resolvers.Workspace.TransferOwnership,
		},
		"createPage": &graphql.Field{
			Type: types.PageType,
			Args: graphql.FieldConfigArgument{
				"workspaceId": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
				"title":       &graphql.ArgumentConfig{Type: graphql.String},
			},
			Resolve: resolvers.Page.CreatePage,
		},
		"updatePage": &graphql.Field{
			Type: types.PageType,
			Args: graphql.FieldConfigArgument{
				"id":    &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
				"title": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
			},
			Resolve: resolvers.Page.UpdatePage,
		},
		"archivePage": &graphql.Field{
			Type: types.PageType,
			Args: graphql.FieldConfigArgument{
				"id":       &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
				"archived": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Boolean)},
			},
			Resolve: resolvers.Page.ArchivePage,
		},
		"deletePageHard": &graphql.Field{
			Type: graphql.Boolean,
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
			},
			Resolve: resolvers.Page.DeletePageHard,
		},
		"createBlock": &graphql.Field{
			Type: types.BlockType,
			Args: graphql.FieldConfigArgument{
				"pageId":        &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
				"parentblockId": &graphql.ArgumentConfig{Type: graphql.String},
				"type":          &graphql.ArgumentConfig{Type: graphql.NewNonNull(types.BlockTypeEnum)},
				"order":         &graphql.ArgumentConfig{Type: graphql.Int},
				"data":          &graphql.ArgumentConfig{Type: types.JSONScalar},
			},
			Resolve: resolvers.Block.CreateBlock,
		},
		"updateBlock": &graphql.Field{
			Type: types.BlockType,
			Args: graphql.FieldConfigArgument{
				"id":   &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
				"type": &graphql.ArgumentConfig{Type: graphql.NewNonNull(types.BlockTypeEnum)},
				"data": &graphql.ArgumentConfig{Type: types.JSONScalar},
			},
			Resolve: resolvers.Block.UpdateBlock,
		},
		"moveBlock": &graphql.Field{
			Type: types.BlockType,
			Args: graphql.FieldConfigArgument{
				"id":            &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
				"pageId":        &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
				"parentblockId": &graphql.ArgumentConfig{Type: graphql.String},
				"order":         &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
			},
			Resolve: resolvers.Block.MoveBlock,
		},
		"deleteBlockTree": &graphql.Field{
			Type: graphql.Boolean,
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
			},
			Resolve: resolvers.Block.DeleteBlockTree,
		},
	},
})
