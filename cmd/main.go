package main

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/onebluesky882/go-http-crud/internal/router"
	"github.com/onebluesky882/go-http-crud/internal/store"
	"github.com/onebluesky882/go-http-crud/pkg/logger"
	"github.com/uptrace/bun"
)

type User struct {
	bun.BaseModel `bun:"table:users,alias:u"`

	ID    int64  `bun:",pk,autoincrement"`
	Name  string `bun:",notnull"`
	Email string `bun:",unique"`
}

func main() {
	log := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{AddSource: true}))
	r := router.New(store.New())
	wrappedRouter := logger.AddLoggerMid(log, logger.LoggerMid(r))

	log.Info("server starting on port : 3008")

	if err := http.ListenAndServe(":3008", wrappedRouter); err != nil {
		log.Error("failed to start server", "error", err)
	}
}
