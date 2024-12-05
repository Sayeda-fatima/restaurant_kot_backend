package controller

import (
	"net/http"
	"strconv"

	"github.com/NazishAhsan/easy_busy_book_laravel/restaurant_kot/model"
	"github.com/NazishAhsan/easy_busy_book_laravel/restaurant_kot/usecase"
	"github.com/labstack/echo/v4"
)

type (
	MenuAllergenController interface {
		GetMenuAllergenList(c echo.Context) error
		CreateMenuAllergen(c echo.Context) error
		UpdateMenuAllergen(c echo.Context) error
		DeleteMenuAllergen(c echo.Context) error
	}

	menuAllergenController struct{
		mu usecase.MenuAllergenUsecase
	}
)

func NewMenuAllergenController(mu usecase.MenuAllergenUsecase) MenuAllergenController{
	return &menuAllergenController{mu}
}

func (ma *menuAllergenController) GetMenuAllergenList(c echo.Context) error{

	menu := c.Param("menuID")
	menuItemID, _ := strconv.Atoi(menu)

	menuAllergenRes, err := ma.mu.GetMenuAllergenList(uint(menuItemID))

	if err != nil{
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, menuAllergenRes)
}

func (ma *menuAllergenController) CreateMenuAllergen(c echo.Context) error{

	menu := c.Param("menuID")
	menuItemID, _ := strconv.Atoi(menu)

	menuAllergen := model.MenuAllergen{}
	if err := c.Bind(&menuAllergen); err != nil{
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	menuAllergen.MenuItemID = uint(menuItemID)
	menuAllergenRes, err := ma.mu.CreateMenuAllergen(menuAllergen)

	if err != nil{
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, menuAllergenRes)
}

func (ma *menuAllergenController) UpdateMenuAllergen(c echo.Context) error{

	menu := c.Param("menuID")
	menuItemID, _ := strconv.Atoi(menu)

	id := c.Param("id")
	menuAllergenID, _ := strconv.Atoi(id)

	menuAllergen := model.MenuAllergen{}
	if err := c.Bind(&menuAllergen); err != nil{
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	menuAllergen.ID = uint(menuAllergenID)
	menuAllergen.MenuItemID = uint(menuItemID)

	menuAllergenRes, err := ma.mu.UpdateMenuAllergen(menuAllergen, uint(menuAllergenID))

	if err != nil{
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, menuAllergenRes)
}

func (ma *menuAllergenController) DeleteMenuAllergen(c echo.Context) error{

	id := c.Param("id")
	menuAllergenID, _ := strconv.Atoi(id)

	menuAllergen := model.MenuAllergen{}
	if err := c.Bind(&menuAllergen); err != nil{
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if err := ma.mu.DeleteMenuAllergen(menuAllergen, uint(menuAllergenID)); err != nil{
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusOK)
}