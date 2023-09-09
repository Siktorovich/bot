package product

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) List() []Product {
	return allPRoducts
}

func (s *Service) Get(idx int) (*Product, error) {
	return &allPRoducts[idx-1], nil
}
