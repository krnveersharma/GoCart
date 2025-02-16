package domain

import "time"

const (
	SELLER = "seller"
	BUYER  = "buyer"
)

type User struct {
	ID        uint      `json:"id" gorm:"PrimaryKey"`
	FirstName string    `json:"firstName" binding:"required"`
	LastName  string    `json:"lastName"`
	Email     string    `json:"email" gorm:"index;unique;not null"`
	Phone     string    `json:"phone" binding:"required"`
	Password  string    `json:"password"`
	Code      int       `json:"code"`
	Expiry    time.Time `json:"expiry"`
	Verified  bool      `json:"verified" gorm:"default:false"`
	UserType  string    `json:"userType" gorm:"default:buyer"`
	CreatedAt time.Time `json:"created_at" gorm:"default:current_timestamp"`
	UpdatedAt time.Time `json:"updated_at" gorm:"default: current_timestamp"`
}
