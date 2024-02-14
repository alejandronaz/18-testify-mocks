package service_test

import (
	"app/app/internal"
	"app/app/internal/repository"
	"app/app/internal/service"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSave(t *testing.T) {

	t.Run("success save with name and rating", func(t *testing.T) {
		// arrange
		// - repository mock
		rp := repository.NewMovieRepositoryMock()
		// declare the Save method should be called, and
		// what should receive and return
		rp.On("Save", &internal.Movie{ID: 0, Name: "Batman", Rating: 8}).Return(nil)
		rp.FuncSave = func(movie *internal.Movie) error {
			movie.ID = 1
			return nil
		}

		// - service
		sv := service.NewMovieService(rp)

		// act
		movie, err := sv.Save("Batman", 8)

		// assert
		require.NoError(t, err)
		expectedMovie := internal.Movie{ID: 1, Name: "Batman", Rating: 8}
		require.Equal(t, expectedMovie, movie)
		// assert that the repository Save method was called once and was called with the expected arguments
		rp.AssertExpectations(t)
	})

	t.Run("error save with empty name", func(t *testing.T) {
		// arrange
		// - repository mock
		rp := repository.NewMovieRepositoryMock()

		// - service
		sv := service.NewMovieService(rp)

		// act
		_, err := sv.Save("", 8)

		// assert
		require.Error(t, err)
		require.ErrorIs(t, err, internal.ErrMovieNotValid)
	})

}
