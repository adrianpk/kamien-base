package api

import (
	"net/http"

	"github.com/adrianpk/kamien"
)

const (
	intServerErrSt = http.StatusInternalServerError
	createdSt      = http.StatusCreated
	noContentSt    = http.StatusNoContent
	okSt           = http.StatusOK
)

var (
	log = kamien.Log
)
