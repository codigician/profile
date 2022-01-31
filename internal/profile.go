package internal

import (
	"time"

	"github.com/codigician/profile/internal/about"
	"github.com/codigician/profile/internal/submission"
)

type CreateProfileReq struct {
	Firstname   string `json:"firstname,"`
	Lastname    string `json:"lastname,"`
	Email       string `json:"email,"`
	PhoneNumber string `json:"phone_number,"`
	Country     string `json:"country,"`
}

func (p *CreateProfileReq) toPersonal() about.Personal {
	return about.Personal{
		Firstname:   p.Firstname,
		Lastname:    p.Lastname,
		Email:       p.Email,
		PhoneNumber: p.PhoneNumber,
		Country:     p.Country,
	}
}

type UpdateProfileReq struct {
	About About `json:"about"`

	Websites []Website `json:"websites"`

	Education   []Education   `json:"education"`
	WorkHistory []WorkHistory `json:"work_history"`
}

func (r *UpdateProfileReq) toAbout() about.About {
	return about.About{
		Headline: r.About.Headline,
		Me:       r.About.Me,
		Personal: about.Personal{
			Firstname:   r.About.Firstname,
			Lastname:    r.About.Lastname,
			Email:       r.About.Email,
			PhoneNumber: r.About.PhoneNumber,
			Country:     r.About.Country,
		},
		Education:   Educations(r.Education).to(),
		WorkHistory: WorkHistories(r.WorkHistory).to(),
		Websites:    Websites(r.Websites).to(),
	}
}

type CreateProfileRes struct {
	ID string `json:"id"`
}

// ProfileRes response of the complete profile information
type ProfileRes struct {
	About About `json:"about"`

	Websites []Website `json:"websites,"`

	Education   []Education   `json:"education,"`
	WorkHistory []WorkHistory `json:"work_history,"`

	Submissions []Submission `json:"submissions"`
}

type About struct {
	Headline    string `json:"headline,"`
	Me          string `json:"me,"`
	Firstname   string `json:"firstname,"`
	Lastname    string `json:"lastname,"`
	Email       string `json:"email,"`
	PhoneNumber string `json:"phone_number,"`
	Country     string `json:"country,"`
}

type Education struct {
	School      string    `json:"school"`
	Program     string    `json:"program"`
	Degree      string    `json:"degree"`
	Description string    `json:"description"`
	Current     bool      `json:"current"`
	StartedAt   time.Time `json:"started_at"`
	EndedAt     time.Time `json:"ended_at"`
}

type Educations []Education

func (e Educations) to() (ed []about.Education) {
	for idx := range e {
		ed = append(ed, about.Education{
			School:      e[idx].School,
			Program:     e[idx].Program,
			Degree:      e[idx].Degree,
			Description: e[idx].Description,
			Current:     e[idx].Current,
			StartedAt:   e[idx].StartedAt,
			EndedAt:     e[idx].EndedAt,
		})
	}
	return ed
}

// time.Time should be changed to string
type WorkHistory struct {
	Company     string    `json:"company"`
	Role        string    `json:"role"`
	Description string    `json:"description"`
	Current     bool      `json:"current"`
	StartedAt   time.Time `json:"started_at"`
	EndedAt     time.Time `json:"ended_at"`
}

type WorkHistories []WorkHistory

func (w WorkHistories) to() (we []about.WorkHistory) {
	for idx := range w {
		we = append(we, about.WorkHistory{
			Company:     w[idx].Company,
			Role:        w[idx].Role,
			Description: w[idx].Description,
			Current:     w[idx].Current,
			StartedAt:   w[idx].StartedAt,
			EndedAt:     w[idx].EndedAt,
		})
	}
	return we
}

type Website struct {
	Title string `json:"title"`
	URL   string `json:"url"`
}

type Websites []Website

func (w Websites) to() (we []about.Website) {
	for idx := range w {
		we = append(we, about.Website{
			Title: w[idx].Title,
			URL:   w[idx].URL,
		})
	}
	return we
}

type Submission struct {
	QuestionID         string    `json:"question_id,"`
	QuestionTitle      string    `json:"question_title,"`
	QuestionDifficulty string    `json:"question_difficulty,"`
	QuestionLink       string    `json:"question_link,"`
	Success            bool      `json:"success,"`
	At                 time.Time `json:"at"`
}

func fromEducation(education []about.Education) (res []Education) {
	for idx := range education {
		res = append(res, Education{
			School:      education[idx].School,
			Program:     education[idx].Program,
			Degree:      education[idx].Degree,
			Description: education[idx].Description,
			Current:     education[idx].Current,
			StartedAt:   education[idx].StartedAt,
			EndedAt:     education[idx].EndedAt,
		})
	}
	return res
}

func fromWorkHistory(workHistory []about.WorkHistory) (res []WorkHistory) {
	for idx := range workHistory {
		res = append(res, WorkHistory{
			Company:     workHistory[idx].Company,
			Role:        workHistory[idx].Role,
			Description: workHistory[idx].Description,
			Current:     workHistory[idx].Current,
			StartedAt:   workHistory[idx].StartedAt,
			EndedAt:     workHistory[idx].EndedAt,
		})
	}
	return res
}

func fromWebsites(websites []about.Website) (res []Website) {
	for idx := range websites {
		res = append(res, Website{
			Title: websites[idx].Title,
			URL:   websites[idx].URL,
		})
	}

	return res
}

func fromSubmissions(submissions []submission.Submission) (res []Submission) {
	for idx := range submissions {
		res = append(res, Submission{
			QuestionID:         submissions[idx].Question.ID,
			QuestionTitle:      submissions[idx].Question.Title,
			QuestionDifficulty: submissions[idx].Question.Difficulty,
			QuestionLink:       submissions[idx].Question.Link,
			Success:            submissions[idx].Success,
			At:                 submissions[idx].At,
		})
	}

	return res
}
