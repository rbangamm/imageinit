package user

import (
	"context"
	"errors"
	"github.com/rbangamm/imageinit/config"
	"github.com/rbangamm/imageinit/repository/user"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	cfg *config.Config
	userRepo *user.Repository
}

func NewService(cfg *config.Config, userRepo *user.Repository) *Service {
	return &Service{
		cfg: cfg,
		userRepo: userRepo,
	}
}

func (s *Service) CreateUser(ctx context.Context, username string, password string) (*user.User, error) {
	filter := bson.M{"uname" : username}
	foundUser, err := s.userRepo.FindUser(ctx, filter)
	if err == nil && foundUser.UserName == username {
		return nil, errors.New("user with username already exists")
	}
	hashedPwd, err := HashPassword(password)
	if err != nil {
		return nil, err
	}
	newUser := &user.User{
		UserName: username,
		Password: hashedPwd,
	}
	createdUser, err := s.userRepo.CreateUser(ctx, newUser)
	if err != nil {
		return nil, err
	}
	return createdUser, err
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}