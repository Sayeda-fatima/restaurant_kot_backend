package repository

import (
	"fmt"

	"github.com/NazishAhsan/easy_busy_book_go/model"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetUserByEmail(user *model.User, email string) error
	GetUserByID(user *model.User, id uint) error
	GetUserCountByOrganization(result *map[string]interface{}, organizationID uint) error
	CreateUser(user *model.User) error
	UpdateUser(user *model.User, jwt string) error
	UpdateUserRefreshToken(user *model.User, jwt string) error
	GetUserByRefreshToken(user *model.User, refreshToken string) error
	CreateResetPasswordToken(user *model.PasswordResetToken) error
	GetUserByToken(user *model.PasswordResetToken, token string) error
	DeleteResetPasswordToken(token *model.PasswordResetToken, email string) error
	ResetPassword(user *model.User, email string, password string) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {

	return &userRepository{db}
}

func (ur *userRepository) GetUserByEmail(user *model.User, email string) error {

	if err := ur.db.Where("email=?", email).First(user).Error; err != nil {
		return err
	}

	return nil

}

func (ur *userRepository) GetUserByID(user *model.User, id uint) error {

	if err := ur.db.Where("id=?", id).First(user).Error; err != nil {
		return err
	}

	return nil
}

func (ur *userRepository) GetUserCountByOrganization(result *map[string]interface{}, organizationID uint) error{

	if err := ur.db.Raw("SELECT IFNULL(count(users.id),0) as total_users, organizations.access_given from organizations left join users on organizations.id=users.organization_id where organizations.id=?", organizationID).Scan(result).Error; err != nil{
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

	result := ur.db.Model(user).Update("api_token", jwt)

	if err := result.Error; err != nil {
		return err
	}

	if result.RowsAffected < 1 {
		return fmt.Errorf("object does not exist")
	}

	return nil
}

func (ur *userRepository) UpdateUserRefreshToken(user *model.User, jwt string) error {

	result := ur.db.Model(user).Update("refresh_token", jwt)

	if err := result.Error; err != nil {
		return err
	}

	if result.RowsAffected < 1 {
		return fmt.Errorf("object does not exist")
	}
	return nil
}

func (ur *userRepository) GetUserByRefreshToken(user *model.User, refreshToken string) error {

	if err := ur.db.Where("refresh_token=?", refreshToken).First(user).Error; err != nil {
		return err
	}
	return nil
}

func (ur *userRepository) CreateResetPasswordToken(user *model.PasswordResetToken) error {

	if err := ur.db.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func (ur *userRepository) GetUserByToken(user *model.PasswordResetToken, token string) error {

	if err := ur.db.Model(user).Where("token=?", token).First(user).Error; err != nil {
		return err
	}
	return nil
}

func (ur *userRepository) DeleteResetPasswordToken(token *model.PasswordResetToken, email string) error {

	result := ur.db.Model(token).Where("email=?", email).Delete(token)
	if err := result.Error; err != nil {
		return err
	}
	return nil
}

func (ur *userRepository) ResetPassword(user *model.User, email string, password string) error {

	result := ur.db.Model(user).Where("email=?", email).Update("password", password)
	if err := result.Error; err != nil {
		return err
	}
	return nil
}
