package usecase

import (
	"os"
	"time"
	"github.com/NazishAhsan/easy_busy_book_go/common"
	"github.com/NazishAhsan/easy_busy_book_go/model"
	"github.com/NazishAhsan/easy_busy_book_go/repository"
	"github.com/NazishAhsan/easy_busy_book_go/validator"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type (
	UserUsecase interface{
		SignUp (user model.User) (model.UserResponse, error)
		Login (user model.User) (string, error)
		Logout (user model.User) (error)
		RefreshToken (refreshToken string) (string, string, error)
	}	

	userUsecase struct{
		ur repository.UserRepository
		uv validator.UserValidator
	}
)


func NewUserUsecase(ur repository.UserRepository, uv validator.UserValidator) UserUsecase {
	return &userUsecase{ur, uv}
}

func (uu *userUsecase) SignUp (user model.User) (model.UserResponse, error) {

	if err := uu.uv.UserValidate(user); err != nil{
		common.Logger.LogError().Msg(err.Error())
		return model.UserResponse{}, err
	} 
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err!=nil{
		common.Logger.LogError().Msg(err.Error())
		return model.UserResponse{}, err
	}

	newUser := model.User{Email: user.Email,Name: user.Name, OrganizationID: user.OrganizationID, Password: string(hash), AccessType: user.AccessType}
	if err := uu.ur.CreateUser(&newUser); err!=nil{
		common.Logger.LogError().Msg(err.Error())
		return model.UserResponse{}, err
	}

	resUser := model.UserResponse{
		ID: newUser.ID,
		Name: newUser.Name,
		Email: newUser.Email,
	}

	return resUser, nil
}

func (uu *userUsecase) Login (user model.User) (string, error){

	if err := uu.uv.UserValidate(user); err!=nil{
		common.Logger.LogError().Msg(err.Error())
		return "", err
	}

	storedUser := model.User{}
	if err := uu.ur.GetUserByEmail(&storedUser, user.Email); err!=nil{
		common.Logger.LogError().Msg(err.Error())
		return "", err
	}

	err := bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(user.Password))
	if err!=nil{
		common.Logger.LogError().Msg(err.Error())
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": storedUser.ID,
		"organization_id": storedUser.OrganizationID,
		"access_type": storedUser.AccessType,
		"exp": time.Now().Add(time.Hour * 100).Unix(),
	})

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": storedUser.ID,
		"organization_id": storedUser.OrganizationID,
		"access_type": storedUser.AccessType,
		"exp": time.Now().Add(time.Hour * 100).Unix(),
	})

	tokenString, err := token.SignedString([]byte (os.Getenv("SECRET")))

	if err!=nil{
		common.Logger.LogError().Msg(err.Error())
		return "", err
	}
	
	refreshTokenString, err := refreshToken.SignedString([]byte(os.Getenv("SECRET_REFRESH")))

	if err != nil{
		common.Logger.LogError().Msg(err.Error())
		return "", err
	}
	// store jwt token to db 
	if err := uu.ur.UpdateUser(&storedUser, tokenString); err!=nil{
		common.Logger.LogError().Msg(err.Error())
		return "", err
	}

	// store refresh token to db
	if err := uu.ur.UpdateUserRefreshToken(&storedUser, refreshTokenString); err != nil{
		return "", err
	}
	return tokenString, nil
}


func (uu *userUsecase) Logout (user model.User) (error){

	storedUser := model.User{}

	if err := uu.ur.GetUserByEmail(&storedUser, user.Email); err!=nil{
		common.Logger.LogError().Msg(err.Error())
		return err
	}
	if err := uu.ur.UpdateUser(&storedUser, ""); err!=nil{
		common.Logger.LogError().Msg(err.Error())
		return err
	}

	return nil
}

func (uu *userUsecase) RefreshToken(refreshToken string) (string, string, error){

	user := model.User{}
	if err := uu.ur.GetUserByRefreshToken(&user, refreshToken); err != nil{
		return "", "", err
	}

	// generate new tokens
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"organization_id": user.OrganizationID,
		"access_type": user.AccessType,
		"exp": time.Now().Add(time.Hour * 100).Unix(),
	})

	newRefreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"organization_id": user.OrganizationID,
		"access_type": user.AccessType,
		"exp": time.Now().Add(time.Hour * 100).Unix(),
	})

	tokenString, err := token.SignedString([]byte (os.Getenv("SECRET")))

	if err!=nil{
		common.Logger.LogError().Msg(err.Error())
		return "", "", err
	}
	
	refreshTokenString, err := newRefreshToken.SignedString([]byte(os.Getenv("SECRET_REFRESH")))

	if err != nil{
		common.Logger.LogError().Msg(err.Error())
		return "", "", err
	}

	// store jwt token to db 
	if err := uu.ur.UpdateUser(&user, tokenString); err!=nil{
		common.Logger.LogError().Msg(err.Error())
		return "", "", err
	}

	// store refresh token to db
	if err := uu.ur.UpdateUserRefreshToken(&user, refreshTokenString); err != nil{
		return "", "", err
	}
	return tokenString, refreshTokenString, nil
}

func (uu *userUsecase) ForgotPassword (user model.User) error{

	storedUser := model.User{}
	if err := uu.ur.GetUserByEmail(&storedUser, user.Email); err != nil{
		return err
	}
	if err := uu.ur.UpdateUserRefreshToken(&storedUser, user.ApiToken); err != nil{
		return err
	}

	return nil
}
