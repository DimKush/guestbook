package UserIn

type UserIn struct {
	Id       int    `json:"id"`
	Username string `json:"username" binding:"required" gorm:"username"`
	Password string `json:"password" binding:"required"`
}
