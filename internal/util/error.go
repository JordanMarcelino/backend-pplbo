package util

import (
	"errors"
	"net/http"
)

var (
	InternalServerError = errors.New("internal server error")
	NotFound            = errors.New("not found")
	BadRequest          = errors.New("bad request")
	Unauthorized        = errors.New("unauthorized")
)

func CheckError(err error) int {
	if errors.Is(err, InternalServerError) {
		return http.StatusInternalServerError
	}
	if errors.Is(err, NotFound) {
		return http.StatusNotFound
	}
	if errors.Is(err, BadRequest) {
		return http.StatusBadRequest
	}

	return http.StatusUnauthorized
}
