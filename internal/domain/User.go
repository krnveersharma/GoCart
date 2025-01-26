package domain

import "time"

type User struct {
	ID        uint      `json:"id"`
	FirstName int       `json:"firstName"`
	LastName  string    `json:"lastName"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	Password  string    `json:"password"`
	Code      int       `json:"code"`
	Expiry    time.Time `json:"expiry"`
	Verified  bool      `json:"verified"`
	UserType  string    `json:"userType"`
}
