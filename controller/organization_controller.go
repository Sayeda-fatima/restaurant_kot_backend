package controller

import (
	"net/http"
	"strconv"

	"github.com/Sayeda-fatima/restaurant_kot_backend/model"
	"github.com/Sayeda-fatima/restaurant_kot_backend/usecase"
	"github.com/labstack/echo/v4"
)

type (
	OrganizationController interface {
		GetOrganizationList(c echo.Context) error
		CreateOrganization(c echo.Context) error
		UpdateOrganization(c echo.Context) error
		DeleteOrganization(c echo.Context) error
	}

	organizationController struct{
		ou usecase.OrganizationUsecase
	}
)

func NewOrganizationController (ou usecase.OrganizationUsecase) OrganizationController{
	return &organizationController{ou}
}

func (oc *organizationController) GetOrganizationList(c echo.Context) error{

	pageNo := c.QueryParam("page")
	page, _ := strconv.Atoi(pageNo)
	organizationRes, err := oc.ou.GetOrganizationList(page)

	if err != nil{
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, organizationRes)
}

func (oc *organizationController) CreateOrganization(c echo.Context) error{

	organization := model.Organization{}

	if err := c.Bind(&organization); err != nil{
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	organizationRes, err := oc.ou.CreateOrganization(organization)

	if err != nil{
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, organizationRes)
}

func (oc *organizationController) UpdateOrganization(c echo.Context) error{

	id := c.Param("id")
	organizationID, _ := strconv.Atoi(id)

	organization := model.Organization{}

	if err := c.Bind(&organization); err != nil{
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	organization.ID = uint(organizationID)
	organizationRes, err := oc.ou.UpdateOrganization(organization, uint(organizationID))
	if err != nil{
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, organizationRes)
}

func (oc *organizationController) DeleteOrganization(c echo.Context) error{

	id := c.Param("id")
	organizationID, _ := strconv.Atoi(id)

	organization := model.Organization{}

	if err := c.Bind(&organization); err != nil{
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if err := oc.ou.DeleteOrganization(organization, uint(organizationID)); err != nil{
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusOK)
}