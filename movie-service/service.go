package main

type MovieService interface {
	GetMovies() ([]*Movie, error)
	GetMovieByID(id int) (*Movie, error)
	CreateMovie(movie *Movie) error
	UpdateMovie(movie *Movie) error
	DeleteMovie(id int) error
}

type movieService struct {
	repo MovieRepository
}

func NewMovieService(repo MovieRepository) MovieService {
	return &movieService{repo}
}

func (s *movieService) GetMovies() ([]*Movie, error) {
	return s.repo.FindAll()
}

func (s *movieService) GetMovieByID(id int) (*Movie, error) {
	return s.repo.FindByID(id)
}

func (s *movieService) CreateMovie(movie *Movie) error {
	return s.repo.Create(movie)
}

func (s *movieService) UpdateMovie(movie *Movie) error {
	return s.repo.Update(movie)
}

func (s *movieService) DeleteMovie(id int) error {
	return s.repo.Delete(id)
}
