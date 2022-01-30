package mongo

import (
	"time"

	"github.com/codigician/profile/internal/about"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type About struct {
	ID          primitive.ObjectID `bson:"_id"`
	Headline    string             `bson:"headline"`
	Me          string             `bson:"me"`
	Personal    Personal           `bson:"personal"`
	Education   []Education        `bson:"education"`
	WorkHistory []WorkHistory      `bson:"work_history"`
	Websites    []Website          `bson:"websites"`
}

type Personal struct {
	Firstname   string `bson:"firstname"`
	Lastname    string `bson:"lastname"`
	Email       string `bson:"email"`
	PhoneNumber string `bson:"phone_number"`
	Country     string `bson:"country"`
}

type Education struct {
	School      string    `bson:"school"`
	Program     string    `bson:"program"`
	Degree      string    `bson:"degree"`
	Description string    `bson:"description"`
	Current     bool      `bson:"current"`
	StartedAt   time.Time `bson:"started_at"`
	EndedAt     time.Time `bson:"ended_at"`
}

type WorkHistory struct {
	Company     string    `bson:"company"`
	Role        string    `bson:"role"`
	Description string    `bson:"description"`
	Current     bool      `bson:"current"`
	StartedAt   time.Time `bson:"started_at"`
	EndedAt     time.Time `bson:"ended_at"`
}

type Website struct {
	Title string `bson:"title"`
	URL   string `bson:"url"`
}

func FromAbout(a *about.About) About {
	return About{
		ID:       primitive.NewObjectID(),
		Headline: a.Headline,
		Me:       a.Me,
		Personal: Personal{
			Firstname:   a.Personal.Firstname,
			Lastname:    a.Personal.Lastname,
			Email:       a.Personal.Email,
			PhoneNumber: a.Personal.PhoneNumber,
			Country:     a.Personal.Country,
		},
		Education:   FromEducation(a.Education),
		WorkHistory: FromWorkHistory(a.WorkHistory),
		Websites:    FromWebsites(a.Websites),
	}
}

func FromEducation(e []about.Education) (education []Education) {
	for idx := range e {
		education = append(education, Education{
			School:      e[idx].School,
			Program:     e[idx].Program,
			Degree:      e[idx].Degree,
			Description: e[idx].Description,
			Current:     e[idx].Current,
			StartedAt:   e[idx].StartedAt,
			EndedAt:     e[idx].EndedAt,
		})
	}

	return education
}

func FromWorkHistory(w []about.WorkHistory) (workHistory []WorkHistory) {
	for idx := range w {
		workHistory = append(workHistory, WorkHistory{
			Company:     w[idx].Company,
			Role:        w[idx].Role,
			Description: w[idx].Description,
			Current:     w[idx].Current,
			StartedAt:   w[idx].StartedAt,
			EndedAt:     w[idx].EndedAt,
		})
	}

	return workHistory
}

func FromWebsites(w []about.Website) (websites []Website) {
	for idx := range w {
		websites = append(websites, Website{
			Title: w[idx].Title,
			URL:   w[idx].URL,
		})
	}

	return websites
}
