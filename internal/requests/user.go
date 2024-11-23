package requests

type LoginRequest struct {
	EmailAddress string `json:"email_address"`
	Password     string `json:"password"`
}
