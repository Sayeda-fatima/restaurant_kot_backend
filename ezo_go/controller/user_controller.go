package controller

import (
	"github.com/NazishAhsan/easy_busy_book_go/model"
	"github.com/NazishAhsan/easy_busy_book_go/usecase"
	"github.com/labstack/echo/v4"
	"net/http"
)

type UserController interface {
	SignUp(c echo.Context) error
	Login(c echo.Context) error
	RefreshToken(c echo.Context) error
	ForgotPassword(c echo.Context) error
	ResetPassword(c echo.Context) error
	Logout(c echo.Context) error
	CsrfToken(c echo.Context) error
}

type userController struct {
	uu usecase.UserUsecase
}

func NewUserController(uu usecase.UserUsecase) UserController {
	return &userController{uu}
}

func (uc *userController) SignUp(c echo.Context) error {
	user := model.User{}
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	userRes, err := uc.uu.SignUp(user)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, userRes)
}

func (uc *userController) Login(c echo.Context) error {
	user := model.User{}

	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	tokenString, err := uc.uu.Login(user)

	if err != nil {
		return c.JSON(http.StatusUnauthorized, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]string{
		"token": tokenString,
	})
}

func (uc *userController) RefreshToken(c echo.Context) error {

	var request struct {
		RefreshToken string `json:"refresh_token" validate:"required"`
	}

	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	accessToken, refreshToken, err := uc.uu.RefreshToken(request.RefreshToken)

	if err != nil {
		return c.JSON(http.StatusUnauthorized, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]string{
		"token":         accessToken,
		"refresh_token": refreshToken,
	})
}

func (uc *userController) ForgotPassword(c echo.Context) error {

	user := model.User{}
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if err := uc.uu.ForgotPassword(user); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, "reset password mail sent")
}

func (uc *userController) ResetPassword(c echo.Context) error {

	user := model.User{}
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	password := c.FormValue("password")
	email := c.FormValue("email")
	token := c.QueryParam("token")

	user.Email = email

	if err := uc.uu.ResetPassword(user, password, token); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, "password reset successfully")
}

func (uc *userController) Logout(c echo.Context) error {

	user := model.User{}

	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if err := uc.uu.Logout(user); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusOK)
}

func (uc *userController) CsrfToken(c echo.Context) error {

	token := c.Get("csrf").(string)
	return c.JSON(http.StatusOK, echo.Map{
		"csrf_token": token,
	})
}
