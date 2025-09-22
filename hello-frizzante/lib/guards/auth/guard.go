package authguard

import (
	"main/lib/auth"
	"main/lib/config"
	"main/lib/core/client"
	"main/lib/core/guard"
	"main/lib/core/receive"
	"main/lib/core/tag"
)

// Tag is the guard tag used to protect routes that require authentication.
const Tag tag.Tag = 1

// New returns an auth guard that validates the Supabase access token cookie.
func New() guard.Guard {
	return guard.Guard{
		Name: "auth",
		Handler: func(c *client.Client, allow func()) {
			cookieName := config.AuthCookieName()
			tok := receive.Cookie(c, cookieName)
			if tok == "" {
				return
			}
			if _, err := auth.ParseAndVerify(tok); err != nil {
				return
			}
			allow()
		},
		Tags: []tag.Tag{Tag},
	}
}
