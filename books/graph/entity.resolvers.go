package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"books/graph/generated"
	"books/graph/model"
	"context"
)

func (r *entityResolver) FindAssignmentByStudentID(ctx context.Context, studentID string) (*model.Assignment, error) {
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

func (r *entityResolver) FindStudentByID(ctx context.Context, id string) (*model.Student, error) {
	var assignments = r.assignments
	var assign *model.Assignment
	if assignments != nil {
		for _, a := range assignments {
			if a.StudentID == id {
				assign = a
			}
		}
	}

	return &model.Student{
		ID: id,
		Assignment: assign,
	}, nil
}

// Entity returns generated.EntityResolver implementation.
func (r *Resolver) Entity() generated.EntityResolver { return &entityResolver{r} }

type entityResolver struct{ *Resolver }
