package service

import "app/app/internal"

func NewMovieService(rp internal.MovieRepository) *MovieService {
	return &MovieService{
		rp: rp,
	}
}

type MovieService struct {
	rp internal.MovieRepository
}

func (s *MovieService) Save(name string, rating int) (internal.Movie, error) {

	// some validations
	if name == "" || rating < 0 || rating > 10 {
		return internal.Movie{}, internal.ErrMovieNotValid
	}

	movie := &internal.Movie{
		ID:     0,
		Name:   name,
		Rating: rating,
	}

	// save movie
	err := s.rp.Save(movie)

	return *movie, err
}

func (s *MovieService) Get(id int) (*internal.Movie, error) {
	return s.rp.Get(id)
}
