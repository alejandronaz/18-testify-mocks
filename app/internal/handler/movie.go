package handler

import (
	"app/app/internal"
	"net/http"
	"strconv"
)

type MovieHandler struct {
	pathp PathParam
	sv    internal.MovieService
}

func NewMovieHandler(pathp PathParam, sv internal.MovieService) *MovieHandler {
	return &MovieHandler{pathp: pathp, sv: sv}
}

func (h *MovieHandler) Get(w http.ResponseWriter, r *http.Request) {
	// get id from path
	id := h.pathp.Read(r, "id")

	// parse id to int
	idInt, err := strconv.Atoi(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("invalid id"))
		return
	}

	// get movie by id
	movie, err := h.sv.Get(idInt)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("movie not found"))
		return
	}

	// write movie to response
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	movieJSON := []byte(
		`{ "id":` + strconv.Itoa(movie.ID) + `,
		"name":"` + movie.Name + `",
		"rating":` + strconv.Itoa(movie.Rating) + `}`,
	)
	w.Write(movieJSON)
}
