package about

type Service struct {
}

func (s *Service) Get() (*About, error) {
	return nil, nil
}

func (s *Service) Update(about About) error {
	return nil
}
