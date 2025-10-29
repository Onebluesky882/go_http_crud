package router

import (
	"net/http"

	"github.com/onebluesky882/go-http-crud/internal/handler"
)

func New(ns handler.NewsStorer) *http.ServeMux {
	r := http.NewServeMux()

	r.HandleFunc("POST /news", handler.PostNews(ns))

	return r
}
