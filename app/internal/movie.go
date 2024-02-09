package internal

import "errors"

type Movie struct {
	ID     int
	Name   string
	Rating int
}

type MovieService interface {
	Save(name string, rating int) (Movie, error)
	Get(id int) (*Movie, error)
}

type MovieRepository interface {
	Save(movie *Movie) error
	Get(id int) (*Movie, error)
}

var (
	ErrMovieNotFound = errors.New("movie not found")
	ErrMovieNotValid = errors.New("movie not valid")
)
