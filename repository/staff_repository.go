package repository

import (
	"fmt"

	"github.com/Sayeda-fatima/restaurant_kot_backend/model"
	"gorm.io/gorm"
)

type (
	StaffRepository interface {
		GetStaffListByOrganization(staff *[]model.Staff, organizationID uint) error
		GetStaffListByRestaurant(staff *[]model.Staff, organizationID uint, restaurantID uint) error
		CreateStaff(staff *model.Staff) error
		UpdateStaff(staff *model.Staff, id uint, organizationID uint, restaurantID uint) error
		DeleteStaff(staff *model.Staff, id uint, organizationID uint, restaurantID uint) error
	}

	staffRepository struct{
		db *gorm.DB
	}
)

func NewStaffRepository (db *gorm.DB) StaffRepository{
	return &staffRepository{db}
}

func (sr *staffRepository) GetStaffListByOrganization(staff *[]model.Staff, organizationID uint) error{

	if err := sr.db.Where("organization_id=? and is_deleted=0", organizationID).Find(staff).Error; err != nil{
		return err
	}
	return nil
}

func (sr *staffRepository) GetStaffListByRestaurant(staff *[]model.Staff, organizationID uint, restaurantID uint) error{

	if err := sr.db.Where("organization_id=? and restaurant_id=? and is_deleted=0", organizationID, restaurantID).Find(staff).Error; err != nil{
		return err
	}
	return nil
}

func (sr *staffRepository) CreateStaff(staff *model.Staff) error{

	if err := sr.db.Create(staff).Error; err != nil{
		return err
	}

	return nil
}

func (sr *staffRepository) UpdateStaff(staff *model.Staff, id uint, organizationID uint, restaurantID uint) error{

	result := sr.db.Model(staff).Where("id=? and organization_id=? and restaurant_id=?", id, organizationID, restaurantID).Updates(staff)

	if err := result.Error; err != nil{
		return err
	}

	if result.RowsAffected < 1 {
		return fmt.Errorf("object does not exist")
	}
	
	return nil
}

func (sr *staffRepository) DeleteStaff(staff *model.Staff, id uint, organizationID uint, restaurantID uint) error{

	result := sr.db.Model(staff).Where("id=? and organization_id=? and restaurant_id=?", id, organizationID, restaurantID).Update("is_deleted", 1)

	if err := result.Error; err != nil{
		return err
	}

	if result.RowsAffected < 1 {
		return fmt.Errorf("object does not exist")
	}
	
	return nil
}