package types

import (
	"github.com/graphql-go/graphql"
)

var WorkspaceType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Workspace",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.String,
		},
		"name": &graphql.Field{
			Type: graphql.String,
		},
		"OwnerID": &graphql.Field{
			Type: graphql.String,
		},
		"Owner": &graphql.Field{
			Type: UserType,
		},
		"Members": &graphql.Field{
			Type: graphql.NewList(WorkspaceMemberType),
		},
		"Pages": &graphql.Field{
			Type: graphql.NewList(PageType),
		},
	},
})
var WorkspaceRoleEnum = graphql.NewEnum(graphql.EnumConfig{
	Name: "WorkspaceRole",
	Values: graphql.EnumValueConfigMap{
		"OWNER":  &graphql.EnumValueConfig{Value: "OWNER"},
		"ADMIN":  &graphql.EnumValueConfig{Value: "ADMIN"},
		"MEMBER": &graphql.EnumValueConfig{Value: "MEMBER"},
	},
})
var WorkspaceMemberType = graphql.NewObject(graphql.ObjectConfig{
	Name: "WorkspaceMember",
	Fields: graphql.Fields{
		"id":          &graphql.Field{Type: graphql.ID},
		"workspaceId": &graphql.Field{Type: graphql.ID},
		"userId":      &graphql.Field{Type: graphql.ID},
		"role":        &graphql.Field{Type: WorkspaceRoleEnum}, // ou enum
		"user": &graphql.Field{
			Type: UserType,
		},
	},
})

