package graph

import "github.com/luiz-from-delta/fc-graphql-with-go/graph/model"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	categories []*model.Category
	courses    []*model.Course
}
