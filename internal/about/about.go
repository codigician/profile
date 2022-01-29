package about

import "time"

type About struct {
	Headline    string
	Me          string
	Personal    Personal
	Education   []Education
	WorkHistory []WorkHistory
	Websites    []Website
}

type Personal struct {
	Firstname   string
	Lastname    string
	Email       string
	PhoneNumber string
	Country     string
}

type Education struct {
	School      string
	Program     string
	Degree      string
	Description string
	Current     bool
	StartedAt   time.Time
	EndedAt     time.Time
}

type WorkHistory struct {
	Company     string
	Role        string
	Description string
	Current     bool
	StartedAt   time.Time
	EndedAt     time.Time
}

type Website struct {
	Title string
	URL   string
}
