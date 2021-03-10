package service

import (
	"go-di-fiber/domain/model"
	"go-di-fiber/repository"
	"go.uber.org/fx"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (us UserService) Register(username, password string) (*model.User, error) {
	return us.repo.Create(username, password)
}

func (us UserService) CheckDuplicate(username string) bool {
	user := us.repo.FindByUsername(username)
	return user.ID != 0
}

var UserServiceModule = fx.Provide(NewUserService)
