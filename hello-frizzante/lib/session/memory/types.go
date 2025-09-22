package session

type Session struct {
	Todos    []Todo    `json:"todos"`
	Lessons  []Lesson  `json:"lessons"`
}

type Todo struct {
	Description string `json:"description"`
	Checked     bool   `json:"checked"`
}

type Lesson struct {
	Student string `json:"student"`
	Date    string `json:"date"`   // YYYY-MM-DD
	Time    string `json:"time"`   // HH:MM
}
