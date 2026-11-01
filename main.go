package main

import (
	"net/http"
	"os"

	"retro/internal/httpserver"
	"retro/internal/logger"
	"retro/internal/repository/memory"
	"retro/internal/usecase"
)

func main() {
	logLevel := os.Getenv("LOG_LEVEL")
	if logLevel == "" {
		logLevel = "info"
	}

	log, closeLog, err := logger.NewLogger(logLevel)
	if err != nil {
		os.Exit(1)
	}
	defer closeLog()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	repo, err := memory.NewStore("data.json")
	if err != nil {
		log.Fatalf("failed to initialize repository: %v", err)
	}

	uc := usecase.NewCatalogUseCase(repo)
	handler := httpserver.NewHandler(uc)
	router := httpserver.SetupRouter(handler, ".")

	log.Infof("server started on port %s (http://localhost:%s)", port, port)

	if err := http.ListenAndServe(":"+port, router); err != nil {
		log.Fatalf("server error: %v", err)
	}
}
