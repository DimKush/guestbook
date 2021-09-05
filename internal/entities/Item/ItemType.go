package Item

type ItemType struct {
	TypeId     int    `json:"type_id" gorm:"type_id"`
	Systemname string `json:"systemname" gorm:"systemname"`
	Fullname   string `json:"fullname" gorm:"fullname"`
}
