package submission

import (
	"context"
	"time"
)

type Repository interface {
	FindAll(ctx context.Context, id string, start, end time.Time) ([]Submission, error)
}

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) FindAllBetween(ctx context.Context, start, end time.Time) ([]Submission, error) {
	return []Submission{
		{
			UserID: "123456",
			Question: Question{
				ID:         "qid-1",
				Title:      "qid-1-title",
				Difficulty: "easy",
				Link:       "qid-1-link",
			},
			Success: false,
			At:      time.Now(),
		},
		{
			UserID: "123456",
			Question: Question{
				ID:         "qid-2",
				Title:      "qid-2-title",
				Difficulty: "hard",
				Link:       "qid-2-link",
			},
			Success: true,
			At:      time.Time{},
		},
	}, nil
}

func (s *Service) ReceiveSubmission(submission *Submission) error {
	return nil
}
