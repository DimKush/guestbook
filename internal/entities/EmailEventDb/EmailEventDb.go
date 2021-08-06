package EmailEventDb

type EmailEventDb struct {
	Id         int    `gorm:"id"`
	Sender     string `gorm:"sender"`
	SenderPass string
	Receiver   string `gorm:"receiver"`
	EmailBody  string `gorm:"email_body"`
}
