package usecase

import (
	"github.com/NazishAhsan/easy_busy_book_laravel/restaurant_kot/model"
	"github.com/NazishAhsan/easy_busy_book_laravel/restaurant_kot/repository"
	"github.com/NazishAhsan/easy_busy_book_laravel/restaurant_kot/validator"
)

type (
	WeeklyStaffScheduleUsecase interface {
		GetWeeklyStaffScheduleList(organizationID uint, restaurantID uint) ([]model.WeeklyStaffScheduleResponse, error)
		CreateWeeklyStaffSchedule(weeklyStaffSchedule model.WeeklyStaffSchedule) (model.WeeklyStaffScheduleResponse, error)
		UpdateWeeklyStaffSchedule(WeeklyStaffSchedule model.WeeklyStaffSchedule, id uint) (model.WeeklyStaffScheduleResponse, error)
		DeleteWeeklyStaffSchedule(WeeklyStaffSchedule model.WeeklyStaffSchedule, id uint) error
	}

	weeklyStaffScheduleUsecase struct{
		wr repository.WeeklyStaffScheduleRepository
		wv validator.WeeklyStaffScheduleValidator
	}
)

func NewWeeklyStaffScheduleUsecase (wr repository.WeeklyStaffScheduleRepository, wv validator.WeeklyStaffScheduleValidator) WeeklyStaffScheduleUsecase{
	return &weeklyStaffScheduleUsecase{wr, wv}
}

func (wu *weeklyStaffScheduleUsecase) GetWeeklyStaffScheduleList(organizationID uint, restaurantID uint) ([]model.WeeklyStaffScheduleResponse, error){

	weeklyStaffSchedule := []model.WeeklyStaffSchedule{}
	if err := wu.wr.GetWeeklyStaffScheduleList(&weeklyStaffSchedule, organizationID, restaurantID); err != nil{
		return nil, err
	}

	resWeeklyStaffSchedule := []model.WeeklyStaffScheduleResponse{}
	for _, v := range(weeklyStaffSchedule){
		res := model.WeeklyStaffScheduleResponse{
			ID: v.ID,
			OrganizationID: v.OrganizationID,
			RestaurantID: v.RestaurantID,
			StaffID: v.StaffID,
			WorkDate: v.WorkDate,
			StartWorkHour: v.StartWorkHour,
			EndWorkHour: v.StartWorkHour,
			IsHoliday: v.IsHoliday,
			IsWeekend: v.IsWeekend,
		}
		resWeeklyStaffSchedule = append(resWeeklyStaffSchedule, res)
	}
	return resWeeklyStaffSchedule, nil
}

func (wu *weeklyStaffScheduleUsecase) CreateWeeklyStaffSchedule(weeklyStaffSchedule model.WeeklyStaffSchedule) (model.WeeklyStaffScheduleResponse, error){

	if err := wu.wv.WeeklyStaffScheduleValidate(weeklyStaffSchedule); err != nil{
		return model.WeeklyStaffScheduleResponse{}, err
	}

	if err := wu.wr.CreateWeeklyStaffSchedule(&weeklyStaffSchedule); err != nil{
		return model.WeeklyStaffScheduleResponse{}, err
	}

	resWeeklyStaffSchedule := model.WeeklyStaffScheduleResponse{
		ID: weeklyStaffSchedule.ID,
		OrganizationID: weeklyStaffSchedule.OrganizationID,
		RestaurantID: weeklyStaffSchedule.RestaurantID,
		StaffID: weeklyStaffSchedule.StaffID,
		WorkDate: weeklyStaffSchedule.WorkDate,
		StartWorkHour: weeklyStaffSchedule.StartWorkHour,
		EndWorkHour: weeklyStaffSchedule.StartWorkHour,
		IsHoliday: weeklyStaffSchedule.IsHoliday,
		IsWeekend: weeklyStaffSchedule.IsWeekend,
	}
	return resWeeklyStaffSchedule, nil
}

func (wu *weeklyStaffScheduleUsecase) UpdateWeeklyStaffSchedule(weeklyStaffSchedule model.WeeklyStaffSchedule, id uint) (model.WeeklyStaffScheduleResponse, error){

	if err := wu.wv.WeeklyStaffScheduleValidate(weeklyStaffSchedule); err != nil{
		return model.WeeklyStaffScheduleResponse{}, err
	}

	if err := wu.wr.UpdateWeeklyStaffSchedule(&weeklyStaffSchedule, id); err != nil{
		return model.WeeklyStaffScheduleResponse{}, err
	}

	resWeeklyStaffSchedule := model.WeeklyStaffScheduleResponse{
		ID: weeklyStaffSchedule.ID,
		OrganizationID: weeklyStaffSchedule.OrganizationID,
		RestaurantID: weeklyStaffSchedule.RestaurantID,
		StaffID: weeklyStaffSchedule.StaffID,
		WorkDate: weeklyStaffSchedule.WorkDate,
		StartWorkHour: weeklyStaffSchedule.StartWorkHour,
		EndWorkHour: weeklyStaffSchedule.StartWorkHour,
		IsHoliday: weeklyStaffSchedule.IsHoliday,
		IsWeekend: weeklyStaffSchedule.IsWeekend,
	}
	return resWeeklyStaffSchedule, nil
}

func (wu *weeklyStaffScheduleUsecase) DeleteWeeklyStaffSchedule(weeklyStaffSchedule model.WeeklyStaffSchedule, id uint) error{

	if err := wu.wr.DeleteWeeklyStaffSchedule(&weeklyStaffSchedule, id); err != nil{
		return err
	}
	return nil
}