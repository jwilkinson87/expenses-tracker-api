package requests

type LoginRequest struct {
	EmailAddress string `json:"email_address" binding:"required,email" validation[required]:"Please provide an email address" validation[email]:"Please provide a valid email address"`
	Password     string `json:"password" binding:"required" validation[required]:"Please provide a password"`
}

type CreateUserRequest struct {
	EmailAddress    string `json:"email_address" binding:"required,email" validation[required]:"Please provide an email address" validation[email]:"Please provide a valid email address"`
	FirstName       string `json:"first_name" binding:"required,alpha" validation[required]:"Please provide a first name" validation[alpha]:"First name must contain only letters"`
	LastName        string `json:"last_name" binding:"required,alpha" validation[required]:"Please provide a last name" validation[alpha]:"Last name must contain only letters"`
	Password        string `json:"password" binding:"required,validpassword" validation[required]:"Please provide a password" validation[validpassword]:"This field requires a valid password. Please enter a password that has at least 1 upper character, 1 lower character, 1 number, 1 special character, and is at least 7 characters in length"`
	ConfirmPassword string `json:"confirm_password" binding:"required,eqfield=Password" validation[required]:"Please confirm your password" validation[eqfield]:"Passwords do not match"`
}

type UpdatePasswordRequest struct {
	CurrentPassword string `json:"current_password" binding:"required" validation[required]:"Please provide your current password"`
	NewPassword     string `json:"new_password" binding:"required,validpassword" validation[required]:"Please provide a new password" validation[validpassword]:"This field requires a valid password. Please enter a password that has at least 1 upper character, 1 lower character, 1 number, 1 special character, and is at least 7 characters in length"`
	ConfirmPassword string `json:"confirm_password" binding:"required,eqfield=NewPassword" validation[required]:"Please confirm your new password" validation[eqfield]:"Passwords do not match"`
}
