package about

import (
	"context"
)

type Repository interface {
	Get(ctx context.Context, id string) (*About, error)
	Save(ctx context.Context, about *About) (string, error)
	Update(ctx context.Context, email string, about *About) error
}

type Service struct {
	repository Repository
}

func NewService(repository Repository) *Service {
	return &Service{repository}
}

func (s *Service) Get(ctx context.Context, id string) (*About, error) {
	return s.repository.Get(ctx, id)
}

func (s *Service) Update(ctx context.Context, email string, about About) error {
	return s.repository.Update(ctx, email, &about)
}

func (s *Service) Create(ctx context.Context, personal Personal) (string, error) {
	about := About{Personal: personal}
	return s.repository.Save(ctx, &about)
}
