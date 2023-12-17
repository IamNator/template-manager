package mailjet

import (
	"context"
	"errors"

	mailgun "github.com/mailgun/mailgun-go/v3"

	"template-manager/email"
)

type Mailjet struct {
	from    string
	mg      *mailgun.MailgunImpl
	domain  string
	apiKeys string
}

func New(domain, apiKeys, from string) *Mailjet {
	mg := mailgun.NewMailgun(domain, apiKeys)
	return &Mailjet{
		from:    from,
		mg:      mg,
		domain:  domain,
		apiKeys: apiKeys,
	}
}

var _ email.Provider = (*Mailjet)(nil)

var templateIDMap = map[email.TemplateID]string{
	email.TemplateIDSignupVerification: "signup_verification",
}

func (m *Mailjet) Send(ctx context.Context, id email.TemplateID, vars map[string]any) error {
	if err := validateVars(vars); err != nil {
		return err
	}
	to := vars["to"].(string)
	subject := vars["subject"].(string)
	return sendTemplateEmail(ctx, m.mg, templateIDMap[id], m.from, to, subject, vars)
}

func validateVars(vars map[string]any) error {
	if _, ok := vars["to"]; !ok {
		return errors.New("missing to")
	}
	if _, ok := vars["subject"]; !ok {
		return errors.New("missing subject")
	}
	return nil
}

func sendTemplateEmail(ctx context.Context, mg *mailgun.MailgunImpl, templateID string, from, to, subject string, variables map[string]any) error {
	message := mg.NewMessage(
		from,
		subject,
		"",
		to,
	)
	message.SetTemplate(templateID)
	if err := message.AddRecipient(to); err != nil {
		return err
	}
	for key, value := range variables {
		if err := message.AddTemplateVariable(key, value); err != nil {
			return err
		}
	}
	_, _, err := mg.Send(ctx, message)
	return err
}
