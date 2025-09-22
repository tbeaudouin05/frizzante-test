package auth

import (
	"strings"

	"main/lib/auth"
	"main/lib/config"
	"main/lib/core/client"
	"main/lib/core/receive"
	"main/lib/core/send"
)

// Session reads an Authorization: Bearer <token> header, verifies it, and sets a HttpOnly cookie.
// Method: POST (also accepts GET for simplicity in dev)
func Session(c *client.Client) {
	// Accept both application/json fetch and form; we only care about Authorization header
	authz := receive.Header(c, "Authorization")
	if !strings.HasPrefix(strings.ToLower(authz), "bearer ") {
		send.Status(c, 400)
		send.Json(c, map[string]string{"error": "missing Bearer token"})
		return
	}
	token := strings.TrimSpace(authz[len("Bearer "):])
	if _, err := auth.ParseAndVerify(token); err != nil {
		send.Status(c, 401)
		send.Json(c, map[string]string{"error": "invalid token"})
		return
	}

	// Set cookie
	name := config.AuthCookieName()
	send.Cookie(c, name, token)
	send.Status(c, 204)
	send.Message(c, "")
}
