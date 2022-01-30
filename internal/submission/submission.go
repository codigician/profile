package submission

import "time"

type Submission struct {
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
