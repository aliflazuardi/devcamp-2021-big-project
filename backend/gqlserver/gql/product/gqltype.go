package Product

import "github.com/graphql-go/graphql"

var ProductType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Product"
	}
)