package httpserver_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"retro/internal/domain"
	"retro/internal/httpserver"
	"retro/internal/repository/memory"
	"retro/internal/usecase"
)

func TestHTTPRouter_GetStateAndCreateItem(t *testing.T) {
	store, err := memory.NewStore("")
	if err != nil {
		t.Fatalf("failed to create store: %v", err)
	}

	uc := usecase.NewCatalogUseCase(store)
	handler := httpserver.NewHandler(uc)
	router := httpserver.SetupRouter(handler, ".")

	req := httptest.NewRequest(http.MethodGet, "/api/data", nil)
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", rec.Code)
	}

	var state domain.AppState
	if err := json.NewDecoder(rec.Body).Decode(&state); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	if len(state.Items) != 12 {
		t.Errorf("expected 12 seed items, got %d", len(state.Items))
	}

	newItem := domain.Item{
		Title:    "Новая модель",
		Category: "models",
		OwnerID:  "user_alex",
	}
	body, _ := json.Marshal(newItem)
	postReq := httptest.NewRequest(http.MethodPost, "/api/items", bytes.NewReader(body))
	postReq.Header.Set("Content-Type", "application/json")
	postRec := httptest.NewRecorder()
	router.ServeHTTP(postRec, postReq)

	if postRec.Code != http.StatusCreated {
		t.Fatalf("expected status 201, got %d", postRec.Code)
	}

	var created domain.Item
	if err := json.NewDecoder(postRec.Body).Decode(&created); err != nil {
		t.Fatalf("failed to decode created item: %v", err)
	}

	if created.Title != "Новая модель" {
		t.Errorf("expected title 'Новая модель', got %s", created.Title)
	}
}
