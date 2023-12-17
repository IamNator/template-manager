package dto

type SignUpRequest struct {
	Email string `json:"email"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LogoutRequest struct {
	Token string `json:"token"`
}

type InitiateResetPasswordRequest struct {
	Email string `json:"email"`
}

type ResetPasswordRequest struct {
	Email       string `json:"email"`
	Token       string `json:"token"`
	NewPassword string `json:"new_password"`
}
