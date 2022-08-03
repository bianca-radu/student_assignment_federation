package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"books/graph/generated"
	"books/graph/model"
	"context"
)

func (r *mutationResolver) CreateAssignment(ctx context.Context, input model.NewAssignment) (*model.Assignment, error) {
	assignment := &model.Assignment{
		Title:     input.Title,
		StudentID: input.StudentID,
	}
	r.assignments = append(r.assignments, assignment)
	return assignment, nil
}

func (r *queryResolver) Assignment(ctx context.Context, studentID string) (*model.Assignment, error) {
	var assignments = r.assignments
	if assignments != nil {
		for _, a := range assignments {
			if a.StudentID == studentID {
				return a, nil
			}
		}
	}

	return nil, nil
}

func (r *queryResolver) Assignments(ctx context.Context) ([]*model.Assignment, error) {
	return r.assignments, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
