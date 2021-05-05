package user

import (
	"context"
	"errors"
	"github.com/rbangamm/imageinit/auth"
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

func (s *Service) CreateUser(ctx context.Context, username string, password string) (string, error) {
	foundUser, err := s.FindUserByUsername(ctx, username)
	if err == nil && foundUser.UserName == username {
		return "", errors.New("user with username already exists")
	}
	hashedPwd, err := HashPassword(password)
	if err != nil {
		return "", err
	}
	newUser := &user.User{
		UserName: username,
		Password: hashedPwd,
	}
	_, err = s.userRepo.CreateUser(ctx, newUser)
	if err != nil {
		return "", err
	}
	token, err := auth.GenerateToken(username)
	if err != nil {
		return "", err
	}
	return token, err
}

func (s *Service) FindUserByUsername(ctx context.Context, username string) (user user.User, err error) {
	filter := bson.M{"uname" : username}
	foundUser, err := s.userRepo.FindUser(ctx, filter)
	if err != nil {
		return user, err
	}
	user = foundUser
	return user, err
}

func (s *Service) LoginUser(ctx context.Context, username string, password string) (string, error) {
	foundUser, err := s.FindUserByUsername(ctx, username)
	if err != nil {
		return "", err
	}
	if !CheckPasswordHash(password, foundUser.Password) {
		return "", errors.New("wrong username or password")
	}
	token, err := auth.GenerateToken(username)
	if err != nil {
		return "", err
	}
	return token, nil
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}