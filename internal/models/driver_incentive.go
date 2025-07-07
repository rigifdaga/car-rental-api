package models

import "time"

type DriverIncentive struct {
    IncentiveID int     `json:"incentive_id" gorm:"primaryKey;column:incentive_id"`
    RentalID    int     `json:"rental_id" gorm:"column:rental_id"`
    Incentive   float64 `json:"incentive" gorm:"column:incentive"`
    CreatedAt   time.Time `json:"created_at" gorm:"column:created_at"`
    UpdatedAt   time.Time `json:"updated_at" gorm:"column:updated_at"`
    
    // Relations
    Rental      Rental  `json:"rental,omitempty" gorm:"foreignKey:RentalID"`
}

func (DriverIncentive) TableName() string {
    return "driver_incentives"
}