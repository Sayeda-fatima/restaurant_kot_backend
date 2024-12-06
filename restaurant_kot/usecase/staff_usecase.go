package usecase

import (
	"github.com/NazishAhsan/easy_busy_book_laravel/restaurant_kot/model"
	"github.com/NazishAhsan/easy_busy_book_laravel/restaurant_kot/repository"
	"github.com/NazishAhsan/easy_busy_book_laravel/restaurant_kot/validator"
)

type (
	StaffUsecase interface {
		GetStaffListByOrganization(organizationID uint) ([]model.StaffResponse, error)
		GetStaffListByRestaurant(organizationID uint, restaurantID uint) ([]model.StaffResponse, error)
		CreateStaff(staff model.Staff) (model.StaffResponse, error)
		UpdateStaff(staff model.Staff, id uint, organizationID uint, restaurantID uint) (model.StaffResponse, error)
		DeleteStaff(staff model.Staff, id uint, organizationID uint, restaurantID uint) error
	}

	staffUsecase struct{
		sr repository.StaffRepository
		sv validator.StaffValidator
	}
)

func NewStaffUsecase (sr repository.StaffRepository, sv validator.StaffValidator) StaffUsecase{
	return &staffUsecase{sr,sv}
}

func (su *staffUsecase) GetStaffListByOrganization(organizationID uint) ([]model.StaffResponse, error){

	staffs := []model.Staff{}
	if err := su.sr.GetStaffListByOrganization(&staffs, organizationID); err != nil{
		return nil, err
	}

	resStaff := []model.StaffResponse{}
	for _, v := range(staffs){
		res := model.StaffResponse{
			ID: v.ID,
			OrganizationID: v.OrganizationID,
			RestaurantID: v.RestaurantID,
			Name: v.Name,
			PhoneNo: v.PhoneNo,
			Role: v.Role,
			Email: v.Email,
		}
		resStaff = append(resStaff, res)
	}

	return resStaff, nil
}

func (su *staffUsecase) GetStaffListByRestaurant(organizationID uint, restaurantID uint) ([]model.StaffResponse, error){

	staffs := []model.Staff{}

	if err := su.sr.GetStaffListByRestaurant(&staffs, organizationID, restaurantID); err != nil{
		return nil, err
	}

	resStaff := []model.StaffResponse{}
	for _, v := range(staffs){
		res := model.StaffResponse{
			ID: v.ID,
			OrganizationID: v.OrganizationID,
			RestaurantID: v.RestaurantID,
			Name: v.Name,
			PhoneNo: v.PhoneNo,
			Role: v.Role,
			Email: v.Email,
		}
		resStaff = append(resStaff, res)
	}

	return resStaff, nil
}

func (su *staffUsecase) CreateStaff(staff model.Staff) (model.StaffResponse, error){

	if err := su.sv.StaffValidate(&staff); err != nil{
		return model.StaffResponse{}, err
	}

	if err := su.sr.CreateStaff(&staff); err != nil{
		return model.StaffResponse{}, err
	}

	resStaff := model.StaffResponse{
		ID: staff.ID,
		OrganizationID: staff.OrganizationID,
		RestaurantID: staff.RestaurantID,
		Name: staff.Name,
		PhoneNo: staff.PhoneNo,
		Email: staff.Email,
		Role: staff.Role,
	}

	return resStaff, nil
}

func (su *staffUsecase) UpdateStaff(staff model.Staff, id uint, organizationID uint, restaurantID uint) (model.StaffResponse, error){

	if err := su.sv.StaffValidate(&staff); err != nil{
		return model.StaffResponse{}, err
	}

	if err := su.sr.UpdateStaff(&staff, id, organizationID, restaurantID); err != nil{
		return model.StaffResponse{}, err
	}

	resStaff := model.StaffResponse{
		ID: staff.ID,
		OrganizationID: staff.OrganizationID,
		RestaurantID: staff.RestaurantID,
		Name: staff.Name,
		Email: staff.Email,
		PhoneNo: staff.PhoneNo,
		Role: staff.Role,
	}

	return resStaff, nil
}

func (su *staffUsecase) DeleteStaff(staff model.Staff, id uint, organizationID uint, restaurantID uint) error{

	if err := su.sr.DeleteStaff(&staff, id, organizationID, restaurantID); err != nil{
		return err
	}
	return nil
}