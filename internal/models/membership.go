package models

import "time"

type Membership struct {
    MembershipID       int     `json:"membership_id" gorm:"primaryKey;column:membership_id"`
    MembershipName     string  `json:"membership_name" gorm:"column:membership_name"`
    DiscountPercentage float64 `json:"discount_percentage" gorm:"column:discount_percentage"`
    CreatedAt          time.Time `json:"created_at" gorm:"column:created_at"`
    UpdatedAt          time.Time `json:"updated_at" gorm:"column:updated_at"`
}

func (Membership) TableName() string {
    return "memberships"
}