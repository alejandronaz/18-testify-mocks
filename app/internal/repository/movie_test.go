package repository_test

import (
	"app/app/internal"
	"app/app/internal/repository"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSave(t *testing.T) {
	// arrange
	// - repository
	rp := repository.NewMovieRepository()
	// - movie
	movie := &internal.Movie{
		Name:   "Batman",
		Rating: 8,
	}

	// act
	err := rp.Save(movie)

	// assert
	require.NoError(t, err)
	require.Equal(t, 1, movie.ID)
}
