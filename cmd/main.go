package main

import (
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/onebluesky882/go-http-crud/internal/router"
	"github.com/uptrace/bun"
)

type User struct {
	bun.BaseModel `bun:"table:users,alias:u"`

	ID    int64  `bun:",pk,autoincrement"`
	Name  string `bun:",notnull"`
	Email string `bun:",unique"`
}

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{AddSource: true}))
	logger.Info("server starting on port :3008")
	r := router.New()
	err := http.ListenAndServe(":3008", r)
	if err != nil {
		log.Fatal(r)
	}
}
