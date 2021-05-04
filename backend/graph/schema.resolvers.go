package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/rbangamm/imageinit/graph/generated"
	"github.com/rbangamm/imageinit/graph/repository"
	"github.com/rbangamm/imageinit/utils"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input repository.NewUser) (*repository.CreateUserResult, error) {
	user, err := r.userService.CreateUser(ctx, input.Username, input.Password)
	res := &repository.CreateUserResult{}
	if err != nil {
		fmt.Printf("%s, %s, %s", input.Username, input.Password, err)
		res.Status = string(utils.RequestStatusFailure)
		res.Error = err.Error()
		return res, nil
	}
	res.ID = user.ID.Hex()
	res.Status = string(utils.RequestStatusSuccess)
	return res, err
}

func (r *mutationResolver) Login(ctx context.Context, input repository.Login) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) RefreshToken(ctx context.Context, input repository.RefreshTokenInput) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
