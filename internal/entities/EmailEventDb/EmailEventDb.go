package EmailEventDb

type EmailEventDb struct {
	Id        int    `gorm:"id"`
	Sender    string `gorm:"sender"`
	Receiver  string `gorm:"receiver"`
	EmailBody string `gorm:"email_body"`
}
