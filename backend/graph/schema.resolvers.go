package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"github.com/rbangamm/imageinit/auth"

	"github.com/rbangamm/imageinit/graph/generated"
	"github.com/rbangamm/imageinit/graph/repository"
	"github.com/rbangamm/imageinit/utils"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input repository.NewUser) (*repository.CreateUserResult, error) {
	token, err := r.userService.CreateUser(ctx, input.Username, input.Password)
	res := &repository.CreateUserResult{}
	if err != nil {
		fmt.Printf("%s, %s", input.Username, err)
		res.Status = string(utils.RequestStatusFailure)
		res.Error = err.Error()
		return res, nil
	}
	res.Token = token
	res.Status = string(utils.RequestStatusSuccess)
	return res, err
}

func (r *mutationResolver) Login(ctx context.Context, input repository.Login) (*repository.LoginResult, error) {
	token, err := r.userService.LoginUser(ctx, input.Username, input.Password)
	res := &repository.LoginResult{}
	if err != nil {
		res.Status = string(utils.RequestStatusFailure)
		res.Error = err.Error()
		return res, nil
	}
	res.Token = token
	res.Status = string(utils.RequestStatusSuccess)
	return res, nil
}

func (r *mutationResolver) RefreshToken(ctx context.Context, input repository.RefreshTokenInput) (*repository.RefreshTokenResult, error) {
	username, err := auth.ParseToken(input.Token)
	res := &repository.RefreshTokenResult{}
	if err != nil {
		res.Status = string(utils.RequestStatusFailure)
		res.Error = err.Error()
		return res, nil
	}
	token, err := auth.GenerateToken(username)
	if err != nil {
		res.Status = string(utils.RequestStatusFailure)
		res.Error = err.Error()
		return res, nil
	}
	res.Token = token
	res.Status = string(utils.RequestStatusSuccess)
	return res, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
