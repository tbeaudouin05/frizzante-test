package auth

import (
	"main/lib/core/client"
	"main/lib/core/send"
	"main/lib/core/view"
)

// Login serves the Login view
func Login(c *client.Client) {
	send.FileOrElse(c, func() {
		send.View(c, view.View{
			Name:  "Login",
			Props: map[string]any{},
		})
	})
}
