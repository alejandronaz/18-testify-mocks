package handler_test

import (
	"app/app/internal"
	"app/app/internal/handler"
	"app/app/internal/repository"
	"app/app/internal/service"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGet(t *testing.T) {
	t.Run("Success get", func(t *testing.T) {
		// arrange

		// - repository mock
		rp := repository.NewMovieRepositoryMock()
		rp.On("Get", 1).Return(&internal.Movie{ID: 1, Name: "Batman", Rating: 8}, nil)

		// - service
		sv := service.NewMovieService(rp)

		// - path param stub
		pathParamStub := handler.PathParamStub{Value: "1"}

		// - handler
		hd := handler.NewMovieHandler(&pathParamStub, sv)

		// request
		req := httptest.NewRequest("GET", "/movies/1", nil)
		// response
		res := httptest.NewRecorder()

		// act
		hd.Get(res, req)

		// assert
		expectedBody := `{"id":1,"name":"Batman","rating":8}`
		require.Equal(t, http.StatusOK, res.Code)
		require.JSONEq(t, expectedBody, res.Body.String())
	})
}
