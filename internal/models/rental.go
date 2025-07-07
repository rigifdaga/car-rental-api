package models

import "time"

type Rental struct {
    RentalID        int       `json:"rental_id" gorm:"primaryKey;column:rental_id"`
    CustomerID      int       `json:"customer_id" gorm:"column:customer_id"`
    CarID           int       `json:"car_id" gorm:"column:car_id"`
    StartRent       time.Time `json:"start_rent" gorm:"column:start_rent;type:date"`
    EndRent         time.Time `json:"end_rent" gorm:"column:end_rent;type:date"`
    TotalCost       float64   `json:"total_cost" gorm:"column:total_cost"`
    Finished        bool      `json:"finished" gorm:"column:finished"`
    Discount        float64   `json:"discount" gorm:"column:discount"`
    BookingTypeID   int       `json:"booking_type_id" gorm:"column:booking_type_id"`
    DriverID        *int      `json:"driver_id,omitempty" gorm:"column:driver_id"`
    TotalDriverCost float64   `json:"total_driver_cost" gorm:"column:total_driver_cost"`
    CreatedAt       time.Time `json:"created_at" gorm:"column:created_at"`
    UpdatedAt       time.Time `json:"updated_at" gorm:"column:updated_at"`
    
    // Relations
    Customer        Customer     `json:"customer,omitempty" gorm:"foreignKey:CustomerID"`
    Car             Car          `json:"car,omitempty" gorm:"foreignKey:CarID"`
    BookingType     BookingType  `json:"booking_type,omitempty" gorm:"foreignKey:BookingTypeID"`
    Driver          *Driver      `json:"driver,omitempty" gorm:"foreignKey:DriverID"`
}

func (Rental) TableName() string {
    return "rentals"
}