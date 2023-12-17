package email

import "context"

type Provider interface {
	Send(ctx context.Context, id, vars map[string]any) error
}
