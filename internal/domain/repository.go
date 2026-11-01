package domain

import "context"

type Repository interface {
	GetState(ctx context.Context) (*AppState, error)
	GetItems(ctx context.Context) ([]Item, error)
	GetItemByID(ctx context.Context, id string) (*Item, error)
	SaveItem(ctx context.Context, item Item) error
	ToggleItemLike(ctx context.Context, itemID, userID string) (*Item, error)
	GetWishes(ctx context.Context) ([]Wish, error)
	SaveWish(ctx context.Context, wish Wish) error
	GetQuestions(ctx context.Context) ([]Question, error)
	SaveQuestion(ctx context.Context, question Question) error
	AddReply(ctx context.Context, questionID string, reply Reply) (*Question, error)
	ToggleHelpful(ctx context.Context, questionID, replyID, authorName string) (*Question, *Recognition, error)
}
