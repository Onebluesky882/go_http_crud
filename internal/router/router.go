package router

import (
	"net/http"

	"github.com/onebluesky882/go-http-crud/internal/handler"
)

func New() *http.ServeMux {
	r := http.NewServeMux()

	r.HandleFunc("GET /", handler.Hello())
	r.HandleFunc("POST /user", handler.CreateUser())
	r.HandleFunc("GET /user", handler.GetAllUser())
	r.HandleFunc("GET /user/{user_id}", handler.GetUserByID())
	r.HandleFunc("PATCH /user/{user_id}", handler.UpdateUser())
	r.HandleFunc("Delete /user/{user_id}", handler.DeleteUser())

	return r
}
