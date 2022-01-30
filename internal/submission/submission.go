package submission

import "time"

type Submission struct {
	UserID   string
	Question Question
	Success  bool
	At       time.Time
}

type Question struct {
	ID         string
	Title      string
	Difficulty string
	Link       string
}
