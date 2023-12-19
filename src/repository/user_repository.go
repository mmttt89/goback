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
		return nil, err
	}
	return &user, nil
}

func (ur *UserRepository) FindByEmail(email string) (*model.User, error) {
	var user model.User
	if err := ur.database.Where("email = ?", email).First(&user).Error; err != nil {
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

func (ur *UserRepository) AddUser(newUser model.User) (*model.User, error) {

	result := ur.database.Create(&newUser)
	if result.Error != nil {
		return nil, result.Error
	}
	return &newUser, nil
}

func (ur *UserRepository) UpdateUser(updatedUser model.User) (bool, error) {
	result := ur.database.Model(&model.User{}).Where("email = ?", updatedUser.Email).Updates(updatedUser)
	if result.Error != nil {
		return false, result.Error
	}
	if result.RowsAffected == 0 {
		return false, errors.New("User not found")
	}
	return true, nil
}
