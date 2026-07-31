package captcha

import (
	"context"
	"flabaxxit/internal/config"

	"github.com/9ssi7/turnstile"
)

type Verifier struct {
	svr turnstile.Service
}

func NewVerifier(config config.Config) *Verifier {
	return &Verifier{svr: turnstile.New(turnstile.Config{Secret: config.Cloudflare_secret})}
}

func (v *Verifier) Verify(ctx context.Context, token, ip string) (bool, error) {
	return v.svr.Verify(ctx, token, ip)
}
