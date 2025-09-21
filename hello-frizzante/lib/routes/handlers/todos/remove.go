package todos

import (
	"strconv"

	"main/lib/core/client"
	"main/lib/core/receive"
	"main/lib/core/send"
	"main/lib/session/memory"
)

func Remove(c *client.Client) {
	s := session.Start(receive.SessionId(c))

	l := int64(len(s.Todos))
	if 0 == l {
		// No index found, ignore the request.
		send.Navigate(c, "/todos")
		return
	}

	is := receive.Query(c, "index")
	if is == "" {
		// No index found, ignore the request.
		send.Navigate(c, "/todos")
		return
	}

	i, e := strconv.ParseInt(is, 10, 64)
	if nil != e {
		send.Navigatef(c, "/todos?error=%s", e.Error())
		return
	}
	if i >= l {
		// Index is out of bounds, ignore the request.
		send.Navigate(c, "/todos")
		return
	}

	s.Todos = append(
		s.Todos[:i],
		s.Todos[i+1:]...,
	)

	send.Navigate(c, "/todos")
}
