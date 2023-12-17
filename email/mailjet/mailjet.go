package mailjet

import (
	"context"

	"template-manager/email"
)

type Mailjet struct {
}

var _ email.Provider = (*Mailjet)(nil)

func (m *Mailjet) Send(ctx context.Context, id, vars map[string]any) error {
	return nil
}
