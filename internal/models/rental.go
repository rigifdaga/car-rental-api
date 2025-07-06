package models

import "time"

type Rental struct {
    RentalID   int       `json:"rental_id" gorm:"primaryKey;column:rental_id"`
    CustomerID int       `json:"customer_id" gorm:"column:customer_id"`
    CarID      int       `json:"car_id" gorm:"column:car_id"`
    StartRent  time.Time `json:"start_rent" gorm:"column:start_rent;type:date"`
    EndRent    time.Time `json:"end_rent" gorm:"column:end_rent;type:date"`
    TotalCost  float64   `json:"total_cost" gorm:"column:total_cost"`
    Finished   bool      `json:"finished" gorm:"column:finished"`
    CreatedAt  time.Time `json:"created_at" gorm:"column:created_at"`
    UpdatedAt  time.Time `json:"updated_at" gorm:"column:updated_at"`
    
    // Relations
    Customer   Customer  `json:"customer,omitempty" gorm:"foreignKey:CustomerID"`
    Car        Car       `json:"car,omitempty" gorm:"foreignKey:CarID"`
}

func (Rental) TableName() string {
    return "rentals"
}