package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fs := http.FileServer(http.Dir("."))
	http.Handle("/", fs)

	fmt.Println("==========================================================")
	fmt.Println("🚀 Сервер «Полки коллекционеров Авито» запущен!")
	fmt.Println("👉 Открой в браузере: http://localhost:" + port)
	fmt.Println("==========================================================")

	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalf("Ошибка при запуске сервера: %v", err)
	}
}
