package repository

import (
	"fmt"

	"github.com/NazishAhsan/easy_busy_book_laravel/restaurant_kot/model"
	"gorm.io/gorm"
)

type (
	WeeklyStaffScheduleRepository interface {
		GetWeeklyStaffScheduleList(weeklyStaffSchedule *[]model.WeeklyStaffSchedule, organizationID uint, restaurantID uint) error
		CreateWeeklyStaffSchedule(weeklyStaffSchedule *model.WeeklyStaffSchedule) error
		UpdateWeeklyStaffSchedule(weeklyStaffSchedule *model.WeeklyStaffSchedule, id uint) error
		DeleteWeeklyStaffSchedule(weeklyStaffSchedule *model.WeeklyStaffSchedule, id uint) error
	}

	weeklyStaffScheduleRepository struct{
		db *gorm.DB
	}
)

func NewWeeklyStaffScheduleRepository (db *gorm.DB) WeeklyStaffScheduleRepository{
	return &weeklyStaffScheduleRepository{db}
}

func (wr *weeklyStaffScheduleRepository) GetWeeklyStaffScheduleList(weeklyStaffSchedule *[]model.WeeklyStaffSchedule, organizationID uint, restaurantID uint) error{

	if err := wr.db.Where("organization_id=? and restaurant_id=?", organizationID, restaurantID).Find(weeklyStaffSchedule).Error; err != nil{
		return err
	}
	return nil
}

func (wr *weeklyStaffScheduleRepository) CreateWeeklyStaffSchedule(weeklyStaffSchedule *model.WeeklyStaffSchedule) error{

	if err := wr.db.Create(weeklyStaffSchedule).Error; err != nil{
		return err
	}
	return nil
}

func (wr *weeklyStaffScheduleRepository) UpdateWeeklyStaffSchedule(weeklyStaffSchedule *model.WeeklyStaffSchedule, id uint) error{

	result := wr.db.Model(weeklyStaffSchedule).Where("id=?", id).Updates(weeklyStaffSchedule)

	if err := result.Error; err != nil{
		return err
	}

	if result.RowsAffected < 1 {
		return fmt.Errorf("object does not exist")
	}
	
	return nil
}

func (wr *weeklyStaffScheduleRepository) DeleteWeeklyStaffSchedule(weeklyStaffSchedule *model.WeeklyStaffSchedule, id uint) error{

	result := wr.db.Model(weeklyStaffSchedule).Where("id=?", id).Delete(weeklyStaffSchedule)

	if err := result.Error; err != nil{
		return err
	}

	if result.RowsAffected < 1 {
		return fmt.Errorf("object does not exist")
	}
	
	return nil
}