package models

import "time"

type BookingType struct {
    BookingTypeID int    `json:"booking_type_id" gorm:"primaryKey;column:booking_type_id"`
    BookingType   string `json:"booking_type" gorm:"column:booking_type"`
    Description   string `json:"description" gorm:"column:description"`
    CreatedAt     time.Time `json:"created_at" gorm:"column:created_at"`
    UpdatedAt     time.Time `json:"updated_at" gorm:"column:updated_at"`
}

func (BookingType) TableName() string {
    return "booking_types"
}