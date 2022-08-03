package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"math/rand"

	"github.com/student/graph/generated"
	"github.com/student/graph/model"
)

func (r *mutationResolver) CreateStudent(ctx context.Context, input model.NewStudent) (*model.Student, error) {
	student := &model.Student{
		Name:   input.Name,
		Grades: input.Grades,
		ID:     fmt.Sprintf("T%d", rand.Int()),
	}
	r.students = append(r.students, student)
	return student, nil
}

func (r *queryResolver) Student(ctx context.Context, name string) (*model.Student, error) {
	var students = r.students
	if students!=nil{
		for _, s := range students {
			if s.Name == name {
				return s,nil
			}
		}
	}

	return nil,nil;
}

func (r *queryResolver) Students(ctx context.Context) ([]*model.Student, error) {
	return r.students, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
