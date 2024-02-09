package repository

import "app/app/internal"

func NewMovieRepository() *MovieRepository {
	return &MovieRepository{}
}

type MovieRepository struct {
	movies []internal.Movie
}

func (r *MovieRepository) Save(movie *internal.Movie) error {
	// get last id
	lastID := len(r.movies)
	// set new id
	movie.ID = lastID + 1
	// append movie to movies
	r.movies = append(r.movies, *movie)
	return nil
}

func (r *MovieRepository) Get(id int) (*internal.Movie, error) {
	// find movie
	for _, movie := range r.movies {
		if movie.ID == id {
			return &movie, nil
		}
	}
	return &internal.Movie{}, internal.ErrMovieNotFound
}
