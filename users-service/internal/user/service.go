package user

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Create(email, password string) (*User, error) {
	u := &User{
		Email:    email,
		Password: password,
	}
	return u, s.repo.Create(u)
}

func (s *Service) Get(id uint) (*User, error) {
	return s.repo.GetByID(id)
}

func (s *Service) List() ([]User, error) {
	return s.repo.List()
}

func (s *Service) Update(id uint, email, password string) (*User, error) {
	return s.repo.Update(id, email, password)
}

func (s *Service) Delete(id uint) error {
	return s.repo.Delete(id)
}
