package User

import "time"

type User struct {
	Id                int       `json:"-" gorm:"column:id"`
	Name              string    `json:"name" gorm:"column:name"`
	Username          string    `json:"username" binding:"required" gorm:"column:username"`
	Email             string    `json:"email" binding:"required" gorm:"column:email"`
	Password          string    `json:"password" binding:"required" gorm:"column:password_hash"`
	Registration_date time.Time `json:"registration_date" binding:"required" gorm:"column:registration_date"`
}
