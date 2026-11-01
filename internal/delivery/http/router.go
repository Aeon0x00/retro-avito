package http

import (
	"net/http"
	"strings"
)

func SetupRouter(h *Handler, staticDir string) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/api/data", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			h.GetState(w, r)
			return
		}
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	})

	mux.HandleFunc("/api/items", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			h.GetItems(w, r)
		case http.MethodPost:
			h.CreateItem(w, r)
		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("/api/items/like", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			h.ToggleLike(w, r)
			return
		}
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	})

	mux.HandleFunc("/api/wishes", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			h.CreateWish(w, r)
			return
		}
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	})

	mux.HandleFunc("/api/questions", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			h.CreateQuestion(w, r)
			return
		}
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	})

	mux.HandleFunc("/api/questions/reply", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			h.AddReply(w, r)
			return
		}
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	})

	mux.HandleFunc("/api/questions/helpful", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			h.ToggleHelpful(w, r)
			return
		}
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	})

	fs := http.FileServer(http.Dir(staticDir))
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if strings.HasPrefix(r.URL.Path, "/api/") {
			http.NotFound(w, r)
			return
		}
		fs.ServeHTTP(w, r)
	})

	return corsMiddleware(mux)
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}
