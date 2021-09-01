package Item

type Item struct {
	Id           int    `json:"id" gorm:"id"`
	ListId       int    `json:"list_id" gorm:"list_id"`
	ItemTypeName string `json:"item_type" gorm:"item_type_name"`
	ItemTypeId   int    `json:"-" gorm:"type_id"`
	Description  string `json:"description" gorm:"description"`
	ListTitle    string `json:"list_title" gorm:"list_title"`
	ItemOwnerId  int    `json:"-"`
}
