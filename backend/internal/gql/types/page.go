package types

import (
	"github.com/graphql-go/graphql"
)
var PageType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Page",
	Fields: graphql.Fields{
		"id":          &graphql.Field{Type: graphql.ID},
		"workspaceId": &graphql.Field{Type: graphql.ID},
		"title":       &graphql.Field{Type: graphql.String},
		"archived":    &graphql.Field{Type: graphql.Boolean},
	},
})
