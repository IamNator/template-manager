package app

import (
	"context"
	"errors"
	"log/slog"
	"time"

	"gorm.io/gorm"

	"template-manager/dto"
	"template-manager/email"
	"template-manager/entity"
)

type App struct {
	Email  email.Provider
	logger *slog.Logger
	db     *gorm.DB // TODO: replace with repository
}

func New(email email.Provider, logger *slog.Logger, db *gorm.DB) *App {
	return &App{
		Email:  email,
		db:     db,
		logger: logger,
	}
}

func (a *App) Signup(ctx context.Context, req dto.SignUpRequest) error {
	var acc = entity.Account{
		Email: req.Email,
	}

	// generate password
	randomPassword := entity.GenerateRandomPassword()
	if err := acc.SetPassword(randomPassword); err != nil {
		a.logger.ErrorContext(ctx, "failed to set password %+v", err)
		return err
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
		"body":    "Your password is " + randomPassword,
	}
	if err := a.Email.Send(ctx, email.TemplateIDSignupVerification, vars); err != nil {
		a.logger.ErrorContext(ctx, "failed to send email %+v", err)
		return err
	}
	return nil
}

func (a *App) Login(ctx context.Context, req dto.LoginRequest) (*dto.LoginResponse, error) {
	// find existing account
	var acc = entity.Account{
		Email: req.Email,
	}
	if err := a.db.Model(&acc).Where("email = ?", req.Email).First(&acc).Error; err != nil {
		return nil, errors.New("account does not exist")
	}
	// check password
	if !acc.ComparePassword(req.Password) {
		return nil, errors.New("invalid password")
	}

	// return token
	return &dto.LoginResponse{
		Token: time.Now().Format(time.RFC3339Nano),
	}, nil
}

func (a *App) Logout(ctx context.Context, req dto.LogoutRequest) error {
	return nil
}

func (a *App) InitiateResetPassword(ctx context.Context, req dto.InitiateResetPasswordRequest) error {
	var acc = entity.Account{
		Email: req.Email,
	}
	if err := a.db.Model(&acc).Where("email = ?", req.Email).First(&acc).Error; err != nil {
		return errors.New("account does not exist")
	}

	// generate password
	randomPassword := entity.GenerateRandomPassword()
	if err := acc.SetPassword(randomPassword); err != nil {
		a.logger.ErrorContext(ctx, "failed to set password %+v", err)
		return err
	}
	//send email
	vars := map[string]any{
		"to":      req.Email,
		"subject": "Welcome to Template Manager",
		"body":    "Your password is " + randomPassword,
	}
	if err := a.Email.Send(ctx, email.TemplateIDSignupVerification, vars); err != nil {
		a.logger.ErrorContext(ctx, "failed to send email %+v", err)
		return err
	}

	return nil
}
