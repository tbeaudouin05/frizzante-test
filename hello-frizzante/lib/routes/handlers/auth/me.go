package auth

import (
	"main/lib/auth"
	"main/lib/config"
	"main/lib/core/client"
	"main/lib/core/receive"
	"main/lib/core/send"
)

// Me returns decoded JWT claims if cookie is valid.
func Me(c *client.Client) {
	name := config.AuthCookieName()
	token := receive.Cookie(c, name)
	if token == "" {
		send.Status(c, 401)
		send.Json(c, map[string]string{"error": "no session"})
		return
	}
	claims, err := auth.ParseAndVerify(token)
	if err != nil {
		send.Status(c, 401)
		send.Json(c, map[string]string{"error": "invalid session"})
		return
	}
	resp := map[string]any{
		"sub":   claims.Subject,
		"email": claims.Email,
		"role":  claims.Role,
		"raw":   claims.Raw,
	}
	send.Json(c, resp)
}
