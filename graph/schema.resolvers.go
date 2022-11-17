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
	id := r.GetNextID()
	newBlog := &model.Blog{
		ID:    id,
		Title: input.Title,
		Body:  input.Body,
		Date:  input.Date,
		Author: &model.Author{
			ID:        input.Author.ID,
			Firstname: input.Author.Firstname,
			Lastname:  input.Author.Lastname,
			Nickname:  input.Author.Nickname,
		},
		Status: model.StatusPublished,
	}
	r.DB[id] = newBlog
	return newBlog, nil
}

// DeleteBlog is the resolver for the deleteBlog field.
func (r *mutationResolver) DeleteBlog(ctx context.Context, id int) (int, error) {
	delete(r.DB, id)
	return id, nil
}

// BlogList is the resolver for the blogList field.
func (r *queryResolver) BlogList(ctx context.Context) ([]*model.Blog, error) {
	var blogList []*model.Blog

	for _, b := range r.DB {
		blogList = append(blogList, b)
	}
	return blogList, nil
}

// GetBlog is the resolver for the getBlog field.
func (r *queryResolver) GetBlog(ctx context.Context, id int) (*model.Blog, error) {
	blog, ok := r.DB[id]
	if !ok {
		return nil, fmt.Errorf("Blog with %d Not found!!", id)
	}
	return blog, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
