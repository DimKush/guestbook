package EventItem

type EventItem struct {
	Id            int    `json:"id" gorm:"id"`
	ListId        int    `json:"list_id" gorm:"list_id"`
	EventTypeName string `json:"event_type" gorm:"event_type_name"`
	EventTypeId   int    `json:"-" gorm:"type_id"`
	Description   string `json:"description" gorm:"description"`
	ListTitle     string `json:"list_title" gorm:"list_title"`
	EventOwnerId  int    `json:"-"`
}
