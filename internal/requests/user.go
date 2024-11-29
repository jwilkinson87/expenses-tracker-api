package requests

type LoginRequest struct {
	EmailAddress string `json:"email_address" binding:"required"`
	Password     string `json:"password" binding:"required"`
}

type CreateUserRequest struct {
	EmailAddress string `json:"email_address" binding:"required,email"`
	FirstName    string `json:"first_name" binding:"required,alpha"`
	LastName     string `json:"last_name" binding:"required,alpha"`
	Password     string `json:"password" binding:"required,matches=^(?=.*[A-Z])(?=.*[0-9])(?=.*[@$!%*?&])[A-Za-z0-9@$!%*?&]{8,}$"`
}
