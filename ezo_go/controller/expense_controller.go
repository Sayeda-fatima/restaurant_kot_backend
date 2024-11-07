package controller

import (
	"net/http"
	"strconv"

	"github.com/NazishAhsan/easy_busy_book_go/model"
	"github.com/NazishAhsan/easy_busy_book_go/usecase"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type (
	ExpenseController interface {
		GetExpenseList(c echo.Context) error
		CreateExpense(c echo.Context) error
		UpdateExpense(c echo.Context) error
		DeleteExpense(c echo.Context) error
		ExpenseReport(c echo.Context) error
	}

	expenseController struct {
		eu usecase.ExpenseUsecase
	}
)

func NewExpenseController(eu usecase.ExpenseUsecase) ExpenseController {
	return &expenseController{eu}
}

func (ec *expenseController) GetExpenseList(c echo.Context) error {

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	organizationID := claims["organization_id"]

	expenseRes, err := ec.eu.GetExpenseList(uint(organizationID.(float64)))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, expenseRes)
}

func (ec *expenseController) CreateExpense(c echo.Context) error {

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	organizationID := claims["organization_id"]

	expense := model.Expense{}

	if err := c.Bind(&expense); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	expense.OrganizationID = uint(organizationID.(float64))
	expense.AmountDue = expense.TotalAmount - expense.AmountPaid
	
	expenseRes, err := ec.eu.CreateExpense(expense)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, expenseRes)
}

func (ec *expenseController) UpdateExpense(c echo.Context) error {

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	organizationID := claims["organization_id"]
	id := c.Param("id")
	expenseID, _ := strconv.Atoi(id)

	expense := model.Expense{}
	if err := c.Bind(&expense); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	expense.OrganizationID = uint(organizationID.(float64))
	expense.ID = uint(expenseID)
	expense.AmountDue = expense.TotalAmount - expense.AmountPaid

	expenseRes, err := ec.eu.UpdateExpense(expense, uint(expenseID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, expenseRes)
}

func (ec *expenseController) DeleteExpense(c echo.Context) error {

	id := c.Param("id")
	expenseID, _ := strconv.Atoi(id)

	expense := model.Expense{}

	if err := c.Bind(&expense); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if err := ec.eu.DeleteExpense(expense, uint(expenseID)); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusOK)
}

func (ec *expenseController) ExpenseReport(c echo.Context) error{

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	organizationID := claims["organization_id"]

	dateFrom := c.FormValue("date_from")
	dateTo := c.FormValue("date_to")
	expenseRes, err := ec.eu.ExpenseReport(uint(organizationID.(float64)), dateFrom, dateTo)

	if err !=nil{
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, expenseRes)
}