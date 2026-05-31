package model

import "time"

type WeeklyStaffSchedule struct {
	ID             uint         `json:"id" gorm:"primaryKey"`
	OrganizationID uint         `json:"organization_id" gorm:"not null"`
	Organization   Organization `json:"-" gorm:"foreignKey:OrganizationID;references:ID" validate:"-"`
	RestaurantID   uint         `json:"restaurant_id" gorm:"not null"`
	Restaurant     Restaurant   `json:"-" gorm:"foreignKey:RestaurantID;references:ID" validate:"-"`
	StaffID        uint         `json:"staff_id" gorm:"not null"`
	Staff          Staff        `json:"-" gorm:"foreignKey:StaffID;references:ID" validate:"-"`
	WorkDate       time.Time    `json:"work_date" gorm:"not null"`
	StartWorkHour  int          `json:"start_work_hour" gorm:"not null;type:int(11)"`
	EndWorkHour    int          `json:"end_work_hour" gorm:"not null;type:int(11)"`
	IsHoliday      bool         `json:"is_holiday" gorm:"not null"`
	IsWeekend      bool         `json:"is_weekend" gorm:"not null"`
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
