package repository

import (
	"fmt"

	"github.com/NazishAhsan/easy_busy_book_go/model"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetUserByEmail(user *model.User, email string) error
	GetUserByID(user *model.User, id uint) error
	CreateUser(user *model.User) error
	UpdateUser (user *model.User, jwt string) error
	UpdateUserRefreshToken (user *model.User, jwt string) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {

	return &userRepository{db}
}

func (ur *userRepository) GetUserByEmail(user *model.User, email string) error {

	if err := ur.db.Where("email=?", email).First(user).Error; err != nil{
		return err
	}

	return nil

}

func (ur *userRepository) GetUserByID(user *model.User, id uint) error{

	if err := ur.db.Where("id=?", id).First(user).Error; err != nil{
		return err
	}

	return nil
}

func (ur *userRepository) CreateUser(user *model.User) error {

	if err := ur.db.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func (ur *userRepository) UpdateUser(user *model.User, jwt string) error {

	result :=ur.db.Model(user).Update("api_token", jwt)

	if err := result.Error; err != nil{
		return err
	}

	if result.RowsAffected < 1{
		return fmt.Errorf("object does not exist")
	}

	return nil
}

func (ur *userRepository) UpdateUserRefreshToken(user *model.User, jwt string) error{

	result := ur.db.Model(user).Update("refresh_token", jwt)

	if err := result.Error; err != nil{
		return err
	}

	if result.RowsAffected < 1{
		return fmt.Errorf("object does not exist")
	}
	return nil
}