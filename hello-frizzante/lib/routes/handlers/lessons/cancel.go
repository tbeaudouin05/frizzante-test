package lessons

import (
    "strconv"

    "main/lib/core/client"
    "main/lib/core/receive"
    "main/lib/core/send"
    session "main/lib/session/memory"
)

func Cancel(c *client.Client) {
    s := session.Start(receive.SessionId(c))

    is := receive.Query(c, "index")
    if is == "" {
        send.Navigate(c, "/lessons")
        return
    }

    i, e := strconv.ParseInt(is, 10, 64)
    if nil != e {
        send.Navigatef(c, "/lessons?error=%s", e.Error())
        return
    }

    l := int64(len(s.Lessons))
    if i >= l || i < 0 {
        send.Navigate(c, "/lessons")
        return
    }

    s.Lessons = append(s.Lessons[:i], s.Lessons[i+1:]...)
    send.Navigate(c, "/lessons")
}
