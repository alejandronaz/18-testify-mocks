package repository

import (
	"app/app/internal"

	"github.com/stretchr/testify/mock"
)

func NewMovieRepositoryMock() *MovieRepositoryMock {
	return &MovieRepositoryMock{}
}

type MovieRepositoryMock struct {
	// mock embebbed
	mock.Mock
	// FuncSave is the mock for the Save method
	FuncSave func(movie *internal.Movie) error
}

func (m *MovieRepositoryMock) Save(movie *internal.Movie) error {

	// i am telling that this method was called with the movie argument,
	// and it should return nil (no error) because i specified that in the test
	outputs := m.Called(movie)

	// call the FuncSave function if it is not nil
	if m.FuncSave != nil {
		m.FuncSave(movie) // to run some process, for example, increment the ID
	}

	// get the error from the first return value (internally it does type assertion)
	err := outputs.Error(0)
	return err
}

func (m *MovieRepositoryMock) Get(id int) (*internal.Movie, error) {
	outputs := m.Called(id)
	// get the movie from the first return value (internally it does type assertion)
	movie := outputs.Get(0).(*internal.Movie)
	err := outputs.Error(1)
	return movie, err
}
