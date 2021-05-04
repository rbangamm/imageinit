package user

import (
	"context"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
)

type Repository struct {
	coll *mgm.Collection
}

func NewRepository() *Repository {
	coll := mgm.Coll(&User{})
	return &Repository{
		coll: coll,
	}
}

func (r *Repository) CreateUser(ctx context.Context, user *User) (*User, error) {
	err := r.coll.CreateWithCtx(ctx, user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *Repository) FindUser(ctx context.Context, filter bson.M) (User, error) {
	user := User{}
	err := r.coll.FirstWithCtx(ctx, filter, &user)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (r *Repository) DeleteUser(ctx context.Context, userID string) error {
	filter := bson.M{}
	filter["_id"] = userID
	_, err := r.coll.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}
	return nil
}

