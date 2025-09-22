package lessons

import (
    "net/url"
    "strings"

    "main/lib/core/client"
    "main/lib/core/receive"
    "main/lib/core/send"
    session "main/lib/session/memory"
)

// Book adds a booking to the current session and navigates back to /lessons.
func Book(c *client.Client) {
    sid := receive.SessionId(c)
    s := session.Start(sid)

    name := strings.TrimSpace(receive.Query(c, "student"))
    date := strings.TrimSpace(receive.Query(c, "date"))
    time := strings.TrimSpace(receive.Query(c, "time"))

    if name == "" || date == "" || time == "" {
        send.Navigatef(c, "/lessons?error=%s", url.QueryEscape("Please provide name, date and time"))
        return
    }

    s.Lessons = append(s.Lessons, session.Lesson{Student: name, Date: date, Time: time})

    send.Navigate(c, "/lessons")
}
