package internal

import (
	"time"

	"github.com/codigician/profile/internal/about"
	"github.com/codigician/profile/internal/submission"
)

type CreateProfileReq struct {
	Firstname   string `json:"firstname,omitempty"`
	Lastname    string `json:"lastname,omitempty"`
	Email       string `json:"email,omitempty"`
	PhoneNumber string `json:"phone_number,omitempty"`
	Country     string `json:"country,omitempty"`
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

type CreateProfileRes struct {
	ID string `json:"id"`
}

// ProfileRes response of the complete profile information
type ProfileRes struct {
	About AboutRes `json:"about"`

	Websites []WebsiteRes `json:"websites,omitempty"`

	Education   []EducationRes   `json:"education,omitempty"`
	WorkHistory []WorkHistoryRes `json:"work_history,omitempty"`

	Submissions []SubmissionRes `json:"submissions"`
}

type AboutRes struct {
	Headline    string `json:"headline,omitempty"`
	Me          string `json:"me,omitempty"`
	Firstname   string `json:"firstname,omitempty"`
	Lastname    string `json:"lastname,omitempty"`
	Email       string `json:"email,omitempty"`
	PhoneNumber string `json:"phone_number,omitempty"`
	Country     string `json:"country,omitempty"`
}

type EducationRes struct {
	School      string    `json:"school,omitempty"`
	Program     string    `json:"program,omitempty"`
	Degree      string    `json:"degree,omitempty"`
	Description string    `json:"description,omitempty"`
	Current     bool      `json:"current,omitempty"`
	StartedAt   time.Time `json:"started_at"`
	EndedAt     time.Time `json:"ended_at"`
}

type WorkHistoryRes struct {
	Company     string    `json:"company"`
	Role        string    `json:"role"`
	Description string    `json:"description"`
	Current     bool      `json:"current"`
	StartedAt   time.Time `json:"started_at"`
	EndedAt     time.Time `json:"ended_at"`
}

type WebsiteRes struct {
	Title string `json:"title"`
	URL   string `json:"url"`
}

type SubmissionRes struct {
	QuestionID         string    `json:"question_id,omitempty"`
	QuestionTitle      string    `json:"question_title,omitempty"`
	QuestionDifficulty string    `json:"question_difficulty,omitempty"`
	QuestionLink       string    `json:"question_link,omitempty"`
	Success            bool      `json:"success,omitempty"`
	At                 time.Time `json:"at"`
}

func fromEducation(education []about.Education) (res []EducationRes) {
	for idx := range education {
		res = append(res, EducationRes{
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

func fromWorkHistory(workHistory []about.WorkHistory) (res []WorkHistoryRes) {
	for idx := range workHistory {
		res = append(res, WorkHistoryRes{
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

func fromWebsites(websites []about.Website) (res []WebsiteRes) {
	for idx := range websites {
		res = append(res, WebsiteRes{
			Title: websites[idx].Title,
			URL:   websites[idx].URL,
		})
	}

	return res
}

func fromSubmissions(submissions []submission.Submission) (res []SubmissionRes) {
	for idx := range submissions {
		res = append(res, SubmissionRes{
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
