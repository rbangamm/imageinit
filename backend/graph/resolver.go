package graph

import "github.com/rbangamm/imageinit/service/user"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	userService *user.Service
}

func NewResolver(userService *user.Service) *Resolver {
	return &Resolver{userService: userService}
}
