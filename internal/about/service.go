package about

import (
	"context"
	"time"
)

type Repository interface {
	Save(ctx context.Context, about *About) error
}

type Service struct {
	repository Repository
}

func NewService(repository Repository) *Service {
	return &Service{repository}
}

func (s *Service) Get(ctx context.Context, id string) (*About, error) {
	return &About{
		Headline: "Software Engineer",
		Me:       "Its about me",
		Personal: Personal{
			Firstname:   "Hikmet Ataberk",
			Lastname:    "Canitez",
			Email:       "hiko@gmail.com",
			PhoneNumber: "+905335555555",
			Country:     "Turkey",
		},
		Education: []Education{
			{
				School:    "Ege University",
				Program:   "Computer Science",
				Degree:    "Bachelor",
				Current:   false,
				StartedAt: time.Time{},
				EndedAt:   time.Time{},
			},
			{
				School:  "Dokuz Eylul University",
				Degree:  "Bachelor",
				Program: "Computer Science",
			},
		},
		WorkHistory: []WorkHistory{
			{
				Company:     "Codigician",
				Role:        "Software Engineer",
				Description: "",
				Current:     true,
			},
		},
		Websites: []Website{
			{
				Title: "Github",
				URL:   "https://github.com/codigician",
			},
		},
	}, nil
}

func (s *Service) Update(ctx context.Context, id string, about About) error {
	return nil
}

func (s *Service) Create(ctx context.Context, personal Personal) error {
	about := About{Personal: personal}
	return s.repository.Save(ctx, &about)
}
