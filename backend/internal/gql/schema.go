package gql

import (
	"github.com/graphql-go/graphql"
	"github.com/ilyes-rhdi/buildit-Gql/internal/gql/mutations"
	"github.com/ilyes-rhdi/buildit-Gql/internal/gql/queries"
)

var Schema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query:    queries.RootQuery,
	Mutation: mutations.RootMutation,
})
