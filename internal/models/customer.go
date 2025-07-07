package models

import "time"

type Customer struct {
    CustomerID   int        `json:"customer_id" gorm:"primaryKey;column:customer_id"`
    Name         string     `json:"name" gorm:"column:name"`
    NIK          string     `json:"nik" gorm:"column:nik;unique"`
    PhoneNumber  string     `json:"phone_number" gorm:"column:phone_number"`
    MembershipID *int       `json:"membership_id,omitempty" gorm:"column:membership_id"`
    CreatedAt    time.Time  `json:"created_at" gorm:"column:created_at"`
    UpdatedAt    time.Time  `json:"updated_at" gorm:"column:updated_at"`
    
    // Relations
    Membership   *Membership `json:"membership,omitempty" gorm:"foreignKey:MembershipID"`
}

func (Customer) TableName() string {
    return "customers"
}