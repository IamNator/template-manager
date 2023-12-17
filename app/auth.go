package app

import (
	"template-manager/dto"
	"template-manager/email"
)

type App struct {
	Email email.Provider
}

func New(email email.Provider) *App {
	return &App{
		Email: email,
	}
}

func (a *App) Signup(req dto.SignUpRequest) error {
	return nil
}

func (a *App) Login(req dto.LoginRequest) error {
	return nil
}

func (a *App) Logout(req dto.LogoutRequest) error {
	return nil
}

func (a *App) InitiateResetPassword(req dto.InitiateResetPasswordRequest) error {
	return nil
}

func (a *App) ResetPassword(req dto.ResetPasswordRequest) error {
	return nil
}
