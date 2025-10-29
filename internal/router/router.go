package router

import (
	"net/http"
)

func New() *http.ServeMux {
	r := http.NewServeMux()

	return r
}
