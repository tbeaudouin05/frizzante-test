package todos

import session "main/lib/session/memory"

type Props struct {
	Todos []session.Todo `json:"todos"`
	Error string         `json:"error"`
}
