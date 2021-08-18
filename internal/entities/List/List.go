package List

type List struct {
	IdStr       string `json:"id,string"`
	Id          int    `json:"id,int",gorm:"id"`
	Title       string `json:"title" gorm:"title"`
	Description string `json:"description" gorm:"description"`
	OwnerId     int    `gorm:"owner_user_id" json:"-"`
	Owner       string `json:"owner"`
}
