package welcome

import (
	"main/lib/core/client"
	"main/lib/core/send"
	"main/lib/core/view"
)

type Props struct {
	Message string `json:"message"`
	Error   string `json:"error"`
}

func View(c *client.Client) {
	send.FileOrElse(c, func() {
		send.View(c, view.View{
			Name: "Welcome",
			Props: Props{
				Message: "hello",
			},
		})
	})
}
