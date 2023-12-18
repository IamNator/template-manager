package app

import (
	"context"
	"errors"
	"log/slog"

	"gorm.io/gorm"

	"template-manager/config"
	"template-manager/dto"
	"template-manager/email"
	"template-manager/entity"
)

type SessionManager interface {
	//create a session
	Create(ctx context.Context, accountID string) (*entity.Session, error)
	//verify a session
	Verify(ctx context.Context, token string) (*entity.Session, error)
	//delete a session
	Expire(ctx context.Context, token string) error
	//delete all sessions for an account
	Delete(ctx context.Context, accountID string) error
}

type App struct {
	config *config.Config
	email  email.Provider
	logger *slog.Logger
	db     *gorm.DB // TODO: replace with repository
	sess   SessionManager
}

func New(config *config.Config, email email.Provider, logger *slog.Logger, db *gorm.DB, sessionManager SessionManager) *App {
	return &App{
		config: config,
		email:  email,
		db:     db,
		logger: logger,
		sess:   sessionManager,
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
		"to":           req.Email,
		"subject":      "Password Setup",
		"password":     randomPassword,
		"company_name": "Template Manager",
	}
	if err := a.email.Send(ctx, email.TemplateIDSignupVerification, vars); err != nil {
		a.logger.ErrorContext(ctx, "failed to send email %+v", err)
		return err
	}
	return nil
}

const (
	LoginFailed = "login failed. please check your email and password and try again"
)

func (a *App) Login(ctx context.Context, req dto.LoginRequest) (*dto.LoginResponse, error) {
	// find existing account
	var acc = entity.Account{
		Email: req.Email,
	}
	if err := a.db.Model(&acc).Where("email = ?", req.Email).First(&acc).Error; err != nil {
		a.logger.InfoContext(ctx, "failed to find account %+v", err)
		return nil, errors.New(LoginFailed)
	}
	// check password
	if !acc.ComparePassword(req.Password) {
		return nil, errors.New(LoginFailed)
	}

	// create session
	sess, err := a.sess.Create(ctx, acc.ID)
	if err != nil {
		a.logger.ErrorContext(ctx, "failed to create session %+v", err)
		return nil, err
	}

	// return token
	return &dto.LoginResponse{
		Email:     acc.Email,
		Token:     sess.Token,
		ExpiresAt: sess.ExpiresAt,
	}, nil
}

func (a *App) Logout(ctx context.Context, req dto.LogoutRequest) error {
	return a.sess.Expire(ctx, req.Token)
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
	if err := a.email.Send(ctx, email.TemplateIDSignupVerification, vars); err != nil {
		a.logger.ErrorContext(ctx, "failed to send email %+v", err)
		return err
	}

	return nil
}
