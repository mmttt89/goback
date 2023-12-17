package repository

import (
	"errors"
	"goback/src/model"

	"gorm.io/gorm"
)

type UserRepository struct {
	database *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		database: db,
	}
}

func (ur *UserRepository) FindByID(userID string) (*model.User, error) {
	// find user from database now
	var user model.User
	if err := ur.database.First(&user, userID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (ur *UserRepository) GetAllUsers() ([]model.User, error) {
	var users []model.User
	if err := ur.database.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func (ur *UserRepository) AddUser(newUser model.User) (model.User, error) {

	result := ur.database.Create(&newUser)
	if result.Error != nil {
		return model.User{}, result.Error
	}
	return newUser, nil
}
