package submission

import (
	"context"
	"time"
)

type Repository struct {
}

type Service struct {
	repo Repository
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) FindAllBetween(ctx context.Context, start, end time.Time) ([]Submission, error) {
	return []Submission{
		{
			Question: Question{
				ID:         "qid",
				Title:      "title",
				Difficulty: "easy",
				Link:       "link",
			},
			Success: false,
			At:      time.Now(),
		},
	}, nil
}

func (s *Service) ReceiveSubmission(submission *Submission) error {
	return nil
}
