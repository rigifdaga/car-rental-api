package models

import "time"

type Car struct {
    CarID     int     `json:"car_id" gorm:"primaryKey;column:car_id"`
    Name      string  `json:"name" gorm:"column:name"`
    Stock     int     `json:"stock" gorm:"column:stock"`
    DailyRent float64 `json:"daily_rent" gorm:"column:daily_rent"`
    CreatedAt time.Time `json:"created_at" gorm:"column:created_at"`
    UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at"`
}

func (Car) TableName() string {
    return "cars"
}