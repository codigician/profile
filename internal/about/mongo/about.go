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

type Educations []Education

type Education struct {
	School      string    `bson:"school"`
	Program     string    `bson:"program"`
	Degree      string    `bson:"degree"`
	Description string    `bson:"description"`
	Current     bool      `bson:"current"`
	StartedAt   time.Time `bson:"started_at"`
	EndedAt     time.Time `bson:"ended_at"`
}

type WorkHistories []WorkHistory

type WorkHistory struct {
	Company     string    `bson:"company"`
	Role        string    `bson:"role"`
	Description string    `bson:"description"`
	Current     bool      `bson:"current"`
	StartedAt   time.Time `bson:"started_at"`
	EndedAt     time.Time `bson:"ended_at"`
}

type Websites []Website

type Website struct {
	Title string `bson:"title"`
	URL   string `bson:"url"`
}

func fromAbout(a *about.About, id primitive.ObjectID) About {
	return About{
		ID:       id,
		Headline: a.Headline,
		Me:       a.Me,
		Personal: Personal{
			Firstname:   a.Personal.Firstname,
			Lastname:    a.Personal.Lastname,
			Email:       a.Personal.Email,
			PhoneNumber: a.Personal.PhoneNumber,
			Country:     a.Personal.Country,
		},
		Education:   fromEducation(a.Education),
		WorkHistory: fromWorkHistory(a.WorkHistory),
		Websites:    fromWebsites(a.Websites),
	}
}

func fromEducation(e []about.Education) (education []Education) {
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

func fromWorkHistory(w []about.WorkHistory) (workHistory []WorkHistory) {
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

func fromWebsites(w []about.Website) (websites []Website) {
	for idx := range w {
		websites = append(websites, Website{
			Title: w[idx].Title,
			URL:   w[idx].URL,
		})
	}

	return websites
}

func (a *About) to() *about.About {
	return &about.About{
		Headline:    a.Headline,
		Me:          a.Me,
		Personal:    a.Personal.to(),
		Education:   Educations(a.Education).to(),
		WorkHistory: WorkHistories(a.WorkHistory).to(),
		Websites:    Websites(a.Websites).to(),
	}
}

func (p Personal) to() about.Personal {
	return about.Personal{
		Firstname:   p.Firstname,
		Lastname:    p.Lastname,
		Email:       p.Email,
		PhoneNumber: p.PhoneNumber,
		Country:     p.Country,
	}
}

func (e Educations) to() []about.Education {
	var education []about.Education
	for idx := range e {
		education = append(education, about.Education{
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

func (w WorkHistories) to() []about.WorkHistory {
	var workHistory []about.WorkHistory
	for idx := range w {
		workHistory = append(workHistory, about.WorkHistory{
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

func (w Websites) to() []about.Website {
	var websites []about.Website
	for idx := range w {
		websites = append(websites, about.Website{
			Title: w[idx].Title,
			URL:   w[idx].URL,
		})
	}

	return websites
}
