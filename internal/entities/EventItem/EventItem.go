package EventItem

type EventItem struct {
	Id            int    `json:"id" gorm:"id"`
	ListId        int    `json:"-" gorm:"list_id"`
	ListTile      string `json:"list_title" gorm:"title"`
	EventTypeName string `json:"event_type" gorm:"systemname"`
	EventTypeId   int    `json:"-" gorm:"type_id"`
	Description   string `json:"description" gorm:"description"`
}
