package fallback

import (
	"main/lib/core/client"
	"main/lib/core/send"
	"main/lib/routes/handlers/welcome"
)

func View(c *client.Client) {
	send.FileOrElse(c, func() { welcome.View(c) })
}
