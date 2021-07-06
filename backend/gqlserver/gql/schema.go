package gql 

import "github.com/graphql-go/graphql"

type SchemaWrapper struct {
	Schema graphql.Schema
}

func NewSchemaWrapper() *SchemaWrapper{
	return &SchemaWrapper{}
}

func (s *SchemaWrapper) int() error {
	schema, err := graphql.NewSchmea(graphql.SchemaConfig{
		query: graphql.NewObject(graphql.ObjectConfig{
			Name: "Query",
			Fields: graphql.Fields{
				
			}
		})
	})
};