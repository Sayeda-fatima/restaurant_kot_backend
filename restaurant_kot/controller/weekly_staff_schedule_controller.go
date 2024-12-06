package controller

import (
	"net/http"
	"strconv"

	"github.com/NazishAhsan/easy_busy_book_laravel/restaurant_kot/model"
	"github.com/NazishAhsan/easy_busy_book_laravel/restaurant_kot/usecase"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type (
	WeeklyStaffScheduleController interface {
		GetWeeklyStaffScheduleList(c echo.Context) error
		CreateWeeklyStaffSchedule(c echo.Context) error
		UpdateWeeklyStaffSchedule(c echo.Context) error
		DeleteWeeklyStaffSchedule(c echo.Context) error
	}

	weeklyStaffScheduleController struct{
		wu usecase.WeeklyStaffScheduleUsecase
	}
)

func NewWeeklyStaffSchedule(wu usecase.WeeklyStaffScheduleUsecase) WeeklyStaffScheduleController{
	return &weeklyStaffScheduleController{wu}
}

func (wc *weeklyStaffScheduleController) GetWeeklyStaffScheduleList(c echo.Context) error{

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	organizationID := claims["organization_id"]
	restaurantID := claims["restaurant_id"]

	weeklyStaffScheduleRes, err := wc.wu.GetWeeklyStaffScheduleList(uint(organizationID.(float64)), uint(restaurantID.(float64)))

	if err != nil{
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, weeklyStaffScheduleRes)
}

func (wc *weeklyStaffScheduleController) CreateWeeklyStaffSchedule(c echo.Context) error{

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	organizationID := claims["organization_id"]
	restaurantID := claims["restaurant_id"]

	weeklyStaffSchedule := model.WeeklyStaffSchedule{}
	if err := c.Bind(&weeklyStaffSchedule); err != nil{
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	weeklyStaffSchedule.OrganizationID = uint(organizationID.(float64))
	weeklyStaffSchedule.RestaurantID = uint(restaurantID.(float64))

	weeklyStaffScheduleRes, err := wc.wu.CreateWeeklyStaffSchedule(weeklyStaffSchedule)

	if err != nil{
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, weeklyStaffScheduleRes)
}

func (wc *weeklyStaffScheduleController) UpdateWeeklyStaffSchedule(c echo.Context) error{

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	organizationID := claims["organization_id"]
	restaurantID := claims["restaurant_id"]

	id := c.Param("id")
	weeklyStaffScheduleID, _ := strconv.Atoi(id)

	weeklyStaffSchedule := model.WeeklyStaffSchedule{}
	if err := c.Bind(&weeklyStaffSchedule); err != nil{
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	weeklyStaffSchedule.ID = uint(weeklyStaffScheduleID)
	weeklyStaffSchedule.OrganizationID = uint(organizationID.(float64))
	weeklyStaffSchedule.RestaurantID = uint(restaurantID.(float64))

	weeklyStaffScheduleRes, err := wc.wu.UpdateWeeklyStaffSchedule(weeklyStaffSchedule, uint(weeklyStaffScheduleID))

	if err != nil{
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, weeklyStaffScheduleRes)
}

func (wc *weeklyStaffScheduleController) DeleteWeeklyStaffSchedule(c echo.Context) error{

	id := c.Param("id")
	weeklyStaffScheduleID, _ := strconv.Atoi(id)

	weeklyStaffSchedule := model.WeeklyStaffSchedule{}
	if err := c.Bind(&weeklyStaffSchedule); err != nil{
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if err := wc.wu.DeleteWeeklyStaffSchedule(weeklyStaffSchedule, uint(weeklyStaffScheduleID)); err != nil{
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusOK)
}