package requests

type LoginRequest struct {
	EmailAddress string `json:"email_address"`
	Password     string `json:"password"`
}

type CreateUserRequest struct {
	EmailAddress string `json:"email_address"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Password     string `json:"password"`
}
