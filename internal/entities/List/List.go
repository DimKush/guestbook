package List

type List struct {
	Id          int    `json:"id" gorm:"id"`
	Title       string `json:"title" gorm:"title"`
	Description string `json:"description" gorm:"description"`
	OwnerId     int    `gorm:"owner_user_id"`
	Owner       string `json:"owner"`
}
