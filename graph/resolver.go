package graph

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.
import (
	"github.com/parikshitg/graphql/graph/model"
)

type Resolver struct {
	DB map[int]*model.Blog
}
