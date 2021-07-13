package main

import (
	"gitlab.com/gitlab-fundamentals/internal/http"
	"go.uber.org/zap"
)

func main() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}

	server := http.NewServer(logger, "8080")
	server.Serve()
}
