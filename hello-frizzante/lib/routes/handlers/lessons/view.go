package lessons

import (
	"main/lib/core/client"
	"main/lib/core/send"
	"main/lib/core/view"
	"main/lib/core/receive"
	"main/lib/session/memory"
)

type Props struct {
	Lessons []session.Lesson `json:"lessons"`
	Error   string           `json:"error"`
}

func View(c *client.Client) {
	s := session.Start(receive.SessionId(c))
	send.View(c, view.View{
		Name:  "Lessons",
		Props: Props{Lessons: s.Lessons, Error: receive.Query(c, "error")},
	})
}
