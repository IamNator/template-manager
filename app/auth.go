package app

import (
	"context"
	"errors"
	"log/slog"

	"template-manager/dto"
	"template-manager/email"
	"template-manager/entity"

	"gorm.io/gorm"
)

type App struct {
	Email  email.Provider
	logger slog.Logger
	db     *gorm.DB // TODO: replace with repository
}

func New(email email.Provider) *App {
	return &App{
		Email: email,
	}
}

func (a *App) Signup(ctx context.Context, req dto.SignUpRequest) error {

	var acc = entity.Account{
		Email: req.Email,
	}

	// find existing account
	if err := a.db.Model(&acc).Where("email = ?", req.Email).First(&acc).Error; err == nil {
		return errors.New("account already exists")
	}
	if err := a.db.Model(&acc).Create(&acc).Error; err != nil {
		a.logger.ErrorContext(ctx, "failed to create account %+v", err)
		return err
	}

	//send email
	vars := map[string]any{
		"to":      req.Email,
		"subject": "Welcome to Template Manager",
	}
	if err := a.Email.Send(ctx, email.TemplateIDSignupVerification, vars); err != nil {
		a.logger.ErrorContext(ctx, "failed to send email %+v", err)
		return err
	}
	return nil
}

func (a *App) Login(ctx context.Context, req dto.LoginRequest) error {
	return nil
}

func (a *App) Logout(ctx context.Context, req dto.LogoutRequest) error {
	return nil
}

func (a *App) InitiateResetPassword(ctx context.Context, req dto.InitiateResetPasswordRequest) error {
	return nil
}

func (a *App) ResetPassword(ctx context.Context, req dto.ResetPasswordRequest) error {
	return nil
}
