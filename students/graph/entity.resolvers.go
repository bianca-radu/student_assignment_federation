package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"github.com/student/graph/generated"
	"github.com/student/graph/model"
)

func (r *entityResolver) FindStudentByID(ctx context.Context, id string) (*model.Student, error) {
	var students = r.students
	if students!=nil{
		for _, s := range students {
			if s.ID == id {
				return s,nil
			}
		}
	}

	return nil, errors.New("Student not found")
}

// Entity returns generated.EntityResolver implementation.
func (r *Resolver) Entity() generated.EntityResolver { return &entityResolver{r} }

type entityResolver struct{ *Resolver }
