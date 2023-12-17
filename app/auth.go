package app

import "template-manager/email"

type App struct {
	Email email.Provider
}

func New(email email.Provider) *App {
	return &App{
		Email: email,
	}
}

func (a *App) Signup() error {

}
