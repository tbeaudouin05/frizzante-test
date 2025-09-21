package todos

import (
	"strconv"

	"main/lib/core/client"
	"main/lib/core/receive"
	"main/lib/core/send"
	"main/lib/session/memory"
)

func Check(c *client.Client) {
	s := session.Start(receive.SessionId(c))

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

	l := int64(len(s.Todos))
	if i >= l {
		// Index is out of bounds, ignore the request.
		send.Navigate(c, "/todos")
		return
	}

	s.Todos[i].Checked = true

	send.Navigate(c, "/todos")
}
