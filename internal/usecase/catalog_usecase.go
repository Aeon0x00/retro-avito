package usecase

import (
	"context"
	"errors"
	"fmt"
	"time"

	"retro/internal/domain"
)

type CatalogUseCase interface {
	GetState(ctx context.Context) (*domain.AppState, error)
	GetItems(ctx context.Context) ([]domain.Item, error)
	CreateItem(ctx context.Context, item domain.Item) (*domain.Item, error)
	ToggleLike(ctx context.Context, itemID, userID string) (*domain.Item, error)
	CreateWish(ctx context.Context, wish domain.Wish) (*domain.Wish, error)
	CreateQuestion(ctx context.Context, question domain.Question) (*domain.Question, error)
	AddReply(ctx context.Context, questionID string, reply domain.Reply) (*domain.Question, error)
	ToggleHelpful(ctx context.Context, questionID, replyID, authorName string) (*domain.Question, *domain.Recognition, error)
}

type catalogUseCase struct {
	repo domain.Repository
}

func NewCatalogUseCase(repo domain.Repository) CatalogUseCase {
	return &catalogUseCase{repo: repo}
}

func (uc *catalogUseCase) GetState(ctx context.Context) (*domain.AppState, error) {
	return uc.repo.GetState(ctx)
}

func (uc *catalogUseCase) GetItems(ctx context.Context) ([]domain.Item, error) {
	return uc.repo.GetItems(ctx)
}

func (uc *catalogUseCase) CreateItem(ctx context.Context, item domain.Item) (*domain.Item, error) {
	if item.Title == "" {
		return nil, errors.New("title is required")
	}
	if item.ID == "" {
		item.ID = fmt.Sprintf("item_%d", time.Now().UnixNano())
	}
	if item.Date == "" {
		item.Date = "Сегодня"
	}
	if item.Likes == nil {
		item.Likes = make([]string, 0)
	}
	if err := uc.repo.SaveItem(ctx, item); err != nil {
		return nil, err
	}
	return &item, nil
}

func (uc *catalogUseCase) ToggleLike(ctx context.Context, itemID, userID string) (*domain.Item, error) {
	if itemID == "" || userID == "" {
		return nil, errors.New("itemId and userId are required")
	}
	return uc.repo.ToggleItemLike(ctx, itemID, userID)
}

func (uc *catalogUseCase) CreateWish(ctx context.Context, wish domain.Wish) (*domain.Wish, error) {
	if wish.Title == "" {
		return nil, errors.New("wish title is required")
	}
	if wish.ID == "" {
		wish.ID = fmt.Sprintf("wish_%d", time.Now().UnixNano())
	}
	if wish.Status == "" {
		wish.Status = "active"
	}
	if wish.Proposals == nil {
		wish.Proposals = make([]domain.WishProposal, 0)
	}
	if err := uc.repo.SaveWish(ctx, wish); err != nil {
		return nil, err
	}
	return &wish, nil
}

func (uc *catalogUseCase) CreateQuestion(ctx context.Context, question domain.Question) (*domain.Question, error) {
	if question.Title == "" {
		return nil, errors.New("question title is required")
	}
	if question.ID == "" {
		question.ID = fmt.Sprintf("q_%d", time.Now().UnixNano())
	}
	if question.Status == "" {
		question.Status = "open"
	}
	if question.CreatedAt == "" {
		question.CreatedAt = time.Now().Format("2006-01-02 15:04")
	}
	if question.Replies == nil {
		question.Replies = make([]domain.Reply, 0)
	}
	if err := uc.repo.SaveQuestion(ctx, question); err != nil {
		return nil, err
	}
	return &question, nil
}

func (uc *catalogUseCase) AddReply(ctx context.Context, questionID string, reply domain.Reply) (*domain.Question, error) {
	if questionID == "" {
		return nil, errors.New("questionId is required")
	}
	if reply.Text == "" {
		return nil, errors.New("reply text is required")
	}
	if reply.ID == "" {
		reply.ID = fmt.Sprintf("rep_%d", time.Now().UnixNano())
	}
	if reply.VersionID == "" {
		reply.VersionID = "v1"
	}
	if reply.CreatedAt == "" {
		reply.CreatedAt = time.Now().Format("2006-01-02 15:04")
	}
	if reply.Clarifications == nil {
		reply.Clarifications = make([]domain.Clarification, 0)
	}
	return uc.repo.AddReply(ctx, questionID, reply)
}

func (uc *catalogUseCase) ToggleHelpful(ctx context.Context, questionID, replyID, authorName string) (*domain.Question, *domain.Recognition, error) {
	if questionID == "" || replyID == "" {
		return nil, nil, errors.New("questionId and replyId are required")
	}
	return uc.repo.ToggleHelpful(ctx, questionID, replyID, authorName)
}
