package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	deliveryHttp "retro/internal/delivery/http"
	"retro/internal/repository/memory"
	"retro/internal/usecase"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	repo, err := memory.NewStore("data.json")
	if err != nil {
		log.Fatalf("failed to initialize repository: %v", err)
	}

	uc := usecase.NewCatalogUseCase(repo)
	handler := deliveryHttp.NewHandler(uc)
	router := deliveryHttp.SetupRouter(handler, ".")

	fmt.Println("==========================================================")
	fmt.Println("🚀 Сервер «Полки коллекционеров Авито» запущен (Clean Architecture)")
	fmt.Println("👉 Открой в браузере: http://localhost:" + port)
	fmt.Println("==========================================================")

	if err := http.ListenAndServe(":"+port, router); err != nil {
		log.Fatalf("server error: %v", err)
	}
}
