package tgram

import (
	"context"

	"github.com/gotd/td/telegram/auth"
	"github.com/gotd/td/tg"
	"golang.org/x/xerrors"
)

// AuthParams contains user authentication params.
type AuthParams struct {
	// Phone number (required).
	Phone string
	// 2FA Password, required only if account has 2FA.
	Password string
	// Function to return code which telegram sents to you
	// for authentication (required).
	CodePrompt func(phone string) (string, error)
	// AllowFlashCall allows phone verification via phone calls.
	AllowFlashCall bool
	// Pass true if the phone number is used on the current device.
	// Ignored if AllowFlashCall is not set.
	CurrentNumber bool
	// If a token that will be included in eventually sent SMSs is required:
	// required in newer versions of android, to use the android SMS receiver APIs.
	AllowAppHash bool
}

// AuthUser performs user authentication.
func (c *Client) AuthUser(ctx context.Context, params AuthParams) error {
	if params.CodePrompt == nil {
		return xerrors.Errorf("code prompt required")
	}

	var authenticator auth.UserAuthenticator
	if params.Password != "" {
		authenticator = auth.Constant(
			params.Phone,
			params.Password,
			auth.CodeAuthenticatorFunc(func(ctx context.Context, sentCode *tg.AuthSentCode) (string, error) {
				return params.CodePrompt(params.Phone)
			}))
	} else {
		authenticator = auth.CodeOnly(
			params.Phone,
			auth.CodeAuthenticatorFunc(func(ctx context.Context, sentCode *tg.AuthSentCode) (string, error) {
				return params.CodePrompt(params.Phone)
			}))
	}

	if err := auth.NewFlow(authenticator, auth.SendCodeOptions{
		AllowFlashCall: params.AllowFlashCall,
		CurrentNumber:  params.CurrentNumber,
		AllowAppHash:   params.AllowAppHash,
	}).Run(ctx, c.telegram.Auth()); err != nil {
		return err
	}

	// TODO: Unnecessary call.
	user, err := c.Self(ctx)
	if err != nil {
		return err
	}

	c.self.Store(user)
	return c.gaps.Auth(c.ctx, c.telegram.API(), user.ID(), user.IsBot(), true)
}

// AuthBot performs bot authentication.
func (c *Client) AuthBot(ctx context.Context, token string) error {
	auth, err := c.telegram.Auth().Bot(ctx, token)
	if err != nil {
		return err
	}

	u, ok := auth.User.(*tg.User)
	if !ok {
		return xerrors.Errorf("unexpected user type: %T", auth.User)
	}

	c.self.Store(u)
	return c.gaps.Auth(c.ctx, c.telegram.API(), u.ID, u.Bot, true)
}

// Self user.
func (c *Client) Self(ctx context.Context) (*User, error) {
	user, err := c.telegram.Self(ctx)
	if err != nil {
		return nil, err
	}

	return &User{
		user:   user,
		client: c,
	}, nil
}
