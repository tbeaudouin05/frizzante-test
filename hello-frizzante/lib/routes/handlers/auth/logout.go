package auth

import (
	"fmt"
	"time"

	"main/lib/config"
	"main/lib/core/client"
	"main/lib/core/send"
)

// Logout clears the auth cookie
func Logout(c *client.Client) {
	name := config.AuthCookieName()
	// Expire the cookie immediately
	expire := time.Now().Add(-time.Hour).UTC().Format(time.RFC1123)
	send.Header(c, "Set-Cookie", fmt.Sprintf("%s=; Path=/; HttpOnly; Expires=%s; Max-Age=0", name, expire))
	// Navigate to home for better UX
	send.Navigate(c, "/")
}
