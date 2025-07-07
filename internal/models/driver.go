package models

import "time"

type Driver struct {
    DriverID    int     `json:"driver_id" gorm:"primaryKey;column:driver_id"`
    Name        string  `json:"name" gorm:"column:name"`
    NIK         string  `json:"nik" gorm:"column:nik;unique"`
    PhoneNumber string  `json:"phone_number" gorm:"column:phone_number"`
    DailyCost   float64 `json:"daily_cost" gorm:"column:daily_cost"`
    CreatedAt   time.Time `json:"created_at" gorm:"column:created_at"`
    UpdatedAt   time.Time `json:"updated_at" gorm:"column:updated_at"`
}

func (Driver) TableName() string {
    return "drivers"
}