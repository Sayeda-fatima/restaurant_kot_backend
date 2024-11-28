package model

import "time"

type WeeklyStaffSchedule struct {
	ID             uint         `json:"id" gorm:"primaryKey"`
	OrganizationID uint         `json:"organization_id"`
	Organization   Organization `json:"-" gorm:"foreignKey:OrganizationID;references:ID"`
	RestaurantID   uint         `json:"restaurant_id"`
	Restaurant     Restaurant   `json:"-" gorm:"foreignKey:RestaurantID;references:ID"`
	StaffID        uint         `json:"staff_id"`
	Staff          Staff        `json:"-" gorm:"foreignKey:StaffID;references:ID"`
	WorkDate       time.Time    `json:"work_date"`
	StartWorkHour  int          `json:"start_work_hour"`
	EndWorkHour    int          `json:"end_work_hour"`
	IsHoliday      bool         `json:"is_holiday"`
	IsWeekend      bool         `json:"is_weekend"`
	CreatedAt      time.Time    `json:"created_at"`
	UpdatedAt      time.Time    `json:"updated_at"`
}

type WeeklyStaffScheduleResponse struct {
	ID             uint      `json:"id"`
	OrganizationID uint      `json:"organization_id"`
	RestaurantID   uint      `json:"restaurant_id"`
	StaffID        uint      `json:"staff_id"`
	WorkDate       time.Time `json:"work_date"`
	StartWorkHour  int       `json:"start_work_hour"`
	EndWorkHour    int       `json:"end_work_hour"`
	IsHoliday      bool      `json:"is_holiday"`
	IsWeekend      bool      `json:"is_weekend"`
}
