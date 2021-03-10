package repository

import (
	"go-di-fiber/domain/model"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (ur *UserRepository) Create(username, password string) (*model.User, error) {
	user := model.CreateNewUser(username, password)
	result := ur.db.Select("Username", "Password").Create(user)
	return user, result.Error
}

func (ur *UserRepository) FindByUsername(username string) *model.User {
	user := &model.User{}
	ur.db.Where("username=?", username).First(user)
	return user
}

var UserRepoModule = fx.Provide(NewUserRepository)
