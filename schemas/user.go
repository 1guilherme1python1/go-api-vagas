package schemas

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID       int
	Email    string
	Password string
}

type UserResponse struct {
	ID       int
	Email    string
	Password string
}
