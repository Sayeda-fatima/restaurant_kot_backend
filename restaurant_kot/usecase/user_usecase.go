package usecase

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/NazishAhsan/easy_busy_book_laravel/restaurant_kot/common"
	"github.com/NazishAhsan/easy_busy_book_laravel/restaurant_kot/model"
	"github.com/NazishAhsan/easy_busy_book_laravel/restaurant_kot/repository"
	"github.com/NazishAhsan/easy_busy_book_laravel/restaurant_kot/validator"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type (
	UserUsecase interface {
		SignUp(user model.User) (model.UserResponse, error)
		Login(user model.User) (string, error)
		Logout(user model.User) error
		RefreshToken(refreshToken string) (string, string, error)
		ForgotPassword(user model.User) error
		ResetPassword(user model.User, password string, token string) error
	}

	userUsecase struct {
		ur repository.UserRepository
		uv validator.UserValidator
		es common.EmailService
	}
)

func NewUserUsecase(ur repository.UserRepository, uv validator.UserValidator, es common.EmailService) UserUsecase {
	return &userUsecase{ur, uv, es}
}

func (uu *userUsecase) SignUp(user model.User) (model.UserResponse, error) {

	if err := uu.uv.UserValidate(user); err != nil {
		common.Logger.LogError().Msg(err.Error())
		return model.UserResponse{}, err
	}

	// check if users can register as per their access given
	var totalUsers map[string]interface{}
	if err := uu.ur.GetUserCountByOrganization(&totalUsers, user.OrganizationID); err != nil {
		return model.UserResponse{}, err
	}
	common.Logger.LogInfo().Msgf("type(total_users): %s, type(access_given): %s", totalUsers["total_users"], totalUsers["access_given"])
	// totalUsers["total_users"] -> type(int64), totalUsers["access_given"] -> type(int32) --> needs conversion for comparison as per current query
	if totalUsers["total_users"].(int64) >= int64(totalUsers["access_given"].(int32)) {
		return model.UserResponse{}, fmt.Errorf("you've reached your limit for signups")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		common.Logger.LogError().Msg(err.Error())
		return model.UserResponse{}, err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":         user.ID,
		"organization_id": user.OrganizationID,
		"restaurant_id":   user.RestaurantID,
		"access_type":     user.AccessType,
		"exp":             time.Now().Add(time.Hour * 100).Unix(),
	})
	tokenString, _ := token.SignedString([]byte(os.Getenv("SECRET")))
	newUser := model.User{Email: user.Email, Name: user.Name, OrganizationID: user.OrganizationID, RestaurantID: user.RestaurantID, Password: string(hash), AccessType: user.AccessType, ApiToken: tokenString}
	if err := uu.ur.CreateUser(&newUser); err != nil {
		common.Logger.LogError().Msg(err.Error())
		return model.UserResponse{}, err
	}

	resUser := model.UserResponse{
		ID:    newUser.ID,
		Name:  newUser.Name,
		Email: newUser.Email,
	}

	return resUser, nil
}

func (uu *userUsecase) Login(user model.User) (string, error) {

	if err := uu.uv.UserValidate(user); err != nil {
		common.Logger.LogError().Msg(err.Error())
		return "", err
	}

	storedUser := model.User{}
	if err := uu.ur.GetUserByEmail(&storedUser, user.Email); err != nil {
		common.Logger.LogError().Msg(err.Error())
		return "", err
	}

	err := bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(user.Password))
	if err != nil {
		common.Logger.LogError().Msg(err.Error())
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":         storedUser.ID,
		"organization_id": storedUser.OrganizationID,
		"restaurant_id":   storedUser.RestaurantID,
		"access_type":     storedUser.AccessType,
		"exp":             time.Now().Add(time.Hour * 100).Unix(),
	})

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":         storedUser.ID,
		"organization_id": storedUser.OrganizationID,
		"restaurant_id":   storedUser.RestaurantID,
		"access_type":     storedUser.AccessType,
		"exp":             time.Now().Add(time.Hour * 100).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		common.Logger.LogError().Msg(err.Error())
		return "", err
	}

	refreshTokenString, err := refreshToken.SignedString([]byte(os.Getenv("SECRET_REFRESH")))

	if err != nil {
		common.Logger.LogError().Msg(err.Error())
		return "", err
	}
	// store jwt token to db
	if err := uu.ur.UpdateUser(&storedUser, tokenString); err != nil {
		common.Logger.LogError().Msg(err.Error())
		return "", err
	}

	// store refresh token to db
	if err := uu.ur.UpdateUserRefreshToken(&storedUser, refreshTokenString); err != nil {
		return "", err
	}
	return tokenString, nil
}

func (uu *userUsecase) Logout(user model.User) error {

	storedUser := model.User{}

	if err := uu.ur.GetUserByEmail(&storedUser, user.Email); err != nil {
		common.Logger.LogError().Msg(err.Error())
		return err
	}
	if err := uu.ur.UpdateUser(&storedUser, ""); err != nil {
		common.Logger.LogError().Msg(err.Error())
		return err
	}

	return nil
}

func (uu *userUsecase) RefreshToken(refreshToken string) (string, string, error) {

	user := model.User{}
	if err := uu.ur.GetUserByRefreshToken(&user, refreshToken); err != nil {
		return "", "", err
	}

	// generate new tokens
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":         user.ID,
		"organization_id": user.OrganizationID,
		"restaurant_id":   user.RestaurantID,
		"access_type":     user.AccessType,
		"exp":             time.Now().Add(time.Hour * 100).Unix(),
	})

	newRefreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":         user.ID,
		"organization_id": user.OrganizationID,
		"restaurant_id":   user.RestaurantID,
		"access_type":     user.AccessType,
		"exp":             time.Now().Add(time.Hour * 100).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		common.Logger.LogError().Msg(err.Error())
		return "", "", err
	}

	refreshTokenString, err := newRefreshToken.SignedString([]byte(os.Getenv("SECRET_REFRESH")))

	if err != nil {
		common.Logger.LogError().Msg(err.Error())
		return "", "", err
	}

	// store jwt token to db
	if err := uu.ur.UpdateUser(&user, tokenString); err != nil {
		common.Logger.LogError().Msg(err.Error())
		return "", "", err
	}

	// store refresh token to db
	if err := uu.ur.UpdateUserRefreshToken(&user, refreshTokenString); err != nil {
		return "", "", err
	}
	return tokenString, refreshTokenString, nil
}

func (uu *userUsecase) ForgotPassword(user model.User) error {

	storedUser := model.User{}
	if err := uu.ur.GetUserByEmail(&storedUser, user.Email); err != nil {
		return err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":           storedUser.Email,
		"organization_id": storedUser.OrganizationID,
		"restaurant_id":   storedUser.RestaurantID,
		"access_type":     storedUser.AccessType,
		"exp":             time.Now().Add(time.Hour * 2).Unix(),
	})

	tokenString, _ := token.SignedString([]byte(os.Getenv("SECRET")))

	if err := uu.ur.CreateResetPasswordToken(&model.PasswordResetToken{Email: storedUser.Email, Token: tokenString}); err != nil {
		return err
	}

	otp := rand.Intn(999999-100000) + 100000
	body := fmt.Sprintf(`otp: %d\ntoken: %s`, otp, tokenString)
	if err := uu.es.SendEmail(storedUser.Email, "Reset Password", body); err != nil {
		return err
	}

	return nil
}

func (uu *userUsecase) ResetPassword(user model.User, password string, token string) error {

	storedUser := model.User{}
	if err := uu.ur.GetUserByEmail(&storedUser, user.Email); err != nil {
		return err
	}

	if err := uu.ur.GetUserByToken(&model.PasswordResetToken{}, token); err != nil {
		return err
	}
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), 10)
	newPassword := string(hash)
	if err := uu.ur.ResetPassword(&user, user.Email, newPassword); err != nil {
		common.Logger.LogError().Fields(map[string]interface{}{"error": err.Error()}).Msg("ResetPassword")
		return err
	}

	if err := uu.ur.DeleteResetPasswordToken(&model.PasswordResetToken{}, storedUser.Email); err != nil {
		common.Logger.LogError().Fields(map[string]interface{}{"error": err.Error()}).Msg("ResetPassword")
		return err
	}
	return nil
}
