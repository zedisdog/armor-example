package requests

type Register struct {
	Mobile               string `json:"mobile" binding:"required"`
	Password             string `json:"password" binding:"required"`
	PasswordConfirmation string `json:"password_confirmation" binding:"required"`
}
