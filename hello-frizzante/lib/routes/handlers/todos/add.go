package todos

import (
	"main/lib/core/client"
	"main/lib/core/receive"
	"main/lib/core/send"
	"main/lib/session/memory"
)

func Add(c *client.Client) {
	s := session.Start(receive.SessionId(c))

	d := receive.Query(c, "description")
	if d == "" {
		send.Navigate(c, "/todos?error=todo description cannot be empty")
		return
	}

	s.Todos = append(s.Todos, session.Todo{
		Checked:     false,
		Description: d,
	})

	send.Navigate(c, "/todos")
}
