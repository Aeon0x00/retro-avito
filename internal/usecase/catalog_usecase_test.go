package usecase_test

import (
	"context"
	"testing"

	"retro/internal/domain"
	"retro/internal/repository/memory"
	"retro/internal/usecase"
)

func TestCatalogUseCase_CreateItemAndToggleLike(t *testing.T) {
	ctx := context.Background()
	store, err := memory.NewStore("")
	if err != nil {
		t.Fatalf("failed to create store: %v", err)
	}

	uc := usecase.NewCatalogUseCase(store)

	newItem := domain.Item{
		Title:    "Test Item",
		Category: "models",
		OwnerID:  "user_alex",
	}

	created, err := uc.CreateItem(ctx, newItem)
	if err != nil {
		t.Fatalf("unexpected error creating item: %v", err)
	}

	if created.ID == "" {
		t.Errorf("expected generated ID, got empty")
	}

	updated, err := uc.ToggleLike(ctx, created.ID, "user_olga")
	if err != nil {
		t.Fatalf("unexpected error toggling like: %v", err)
	}

	if len(updated.Likes) != 1 || updated.Likes[0] != "user_olga" {
		t.Errorf("expected likes to contain user_olga, got %v", updated.Likes)
	}

	updatedAgain, err := uc.ToggleLike(ctx, created.ID, "user_olga")
	if err != nil {
		t.Fatalf("unexpected error toggling like again: %v", err)
	}

	if len(updatedAgain.Likes) != 0 {
		t.Errorf("expected likes to be empty, got %v", updatedAgain.Likes)
	}
}
