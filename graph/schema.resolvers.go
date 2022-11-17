package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/parikshitg/graphql/graph/generated"
	"github.com/parikshitg/graphql/graph/model"
)

// CreateBlog is the resolver for the createBlog field.
func (r *mutationResolver) CreateBlog(ctx context.Context, input model.BlogInput) (*model.Blog, error) {
	panic(fmt.Errorf("not implemented: CreateBlog - createBlog"))
}

// DeleteBlog is the resolver for the deleteBlog field.
func (r *mutationResolver) DeleteBlog(ctx context.Context, id int) (int, error) {
	panic(fmt.Errorf("not implemented: DeleteBlog - deleteBlog"))
}

// BlogList is the resolver for the blogList field.
func (r *queryResolver) BlogList(ctx context.Context) ([]*model.Blog, error) {
	panic(fmt.Errorf("not implemented: BlogList - blogList"))
}

// GetBlog is the resolver for the getBlog field.
func (r *queryResolver) GetBlog(ctx context.Context, id int) (*model.Blog, error) {
	panic(fmt.Errorf("not implemented: GetBlog - getBlog"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
