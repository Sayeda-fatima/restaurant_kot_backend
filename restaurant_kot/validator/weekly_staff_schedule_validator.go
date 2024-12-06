package validator

import (
	"github.com/NazishAhsan/easy_busy_book_laravel/restaurant_kot/model"
	"github.com/go-playground/validator"
)

type (
	WeeklyStaffScheduleValidator interface {
		WeeklyStaffScheduleValidate(weeklyStaffSchedule model.WeeklyStaffSchedule) error
	}

	weeklyStaffScheduleValidator struct{
		validator *validator.Validate
	}
)

func NewWeeklyStaffScheduleValidator() WeeklyStaffScheduleValidator{
	return &weeklyStaffScheduleValidator{
		validator: validator.New(),
	}
}

func (wr *weeklyStaffScheduleValidator) WeeklyStaffScheduleValidate(weeklyStaffSchedule model.WeeklyStaffSchedule) error{

	if err := wr.validator.Struct(weeklyStaffSchedule); err != nil{
		return err
	}
	return nil
}