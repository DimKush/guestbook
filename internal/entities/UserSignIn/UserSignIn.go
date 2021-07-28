package UserSighnIn

type UserSignIn struct {
	Username string `json:"username" binding:"required"`
	Password string
}
