package dto

type SignUpRequest struct {
	Email string `json:"email"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LogoutRequest struct {
	AccountID string `json:"account_id"`
	Token     string `json:"token"`
}

type InitiateResetPasswordRequest struct {
	Email string `json:"email"`
}

type ResetPasswordRequest struct {
	Email       string `json:"email"`
	Token       string `json:"token"`
	NewPassword string `json:"new_password"`
}

type DeleteAccessKeyRequest struct {
	AccountID   string `json:"account_id"`
	AccessKeyID string `json:"access_key_id"`
}

type ListAccessKeysRequest struct {
	AccountID string `json:"account_id"`
}

type CreateAccessKeyRequest struct {
	AccountID     string `json:"account_id"`
	AccessKeyName string `json:"access_key_name"`
}
