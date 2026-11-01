package memory

import (
	"context"
	_ "embed"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"sync"
	"time"

	"retro/internal/domain"
)

//go:embed seed.json
var defaultSeed []byte

type Store struct {
	mu           sync.RWMutex
	dataFile     string
	categories   []domain.Category
	items        []domain.Item
	wishes       []domain.Wish
	questions    []domain.Question
	recognitions []domain.Recognition
}

func NewStore(dataFile string) (*Store, error) {
	s := &Store{
		dataFile: dataFile,
	}

	loaded := false
	if dataFile != "" {
		if data, err := os.ReadFile(dataFile); err == nil && len(data) > 0 {
			var state domain.AppState
			if err := json.Unmarshal(data, &state); err == nil {
				s.categories = state.Categories
				s.items = state.Items
				s.wishes = state.Wishes
				s.questions = state.Questions
				s.recognitions = state.Recognitions
				loaded = true
			}
		}
	}

	if !loaded {
		var seed domain.AppState
		if err := json.Unmarshal(defaultSeed, &seed); err != nil {
			return nil, fmt.Errorf("failed to parse seed data: %w", err)
		}
		s.categories = seed.Categories
		s.items = seed.Items
		s.wishes = seed.Wishes
		s.questions = seed.Questions
		s.recognitions = seed.Recognitions
		_ = s.persist()
	}

	return s, nil
}

func (s *Store) persist() error {
	if s.dataFile == "" {
		return nil
	}
	state := domain.AppState{
		Categories:   s.categories,
		Items:        s.items,
		Wishes:       s.wishes,
		Questions:    s.questions,
		Recognitions: s.recognitions,
	}
	data, err := json.MarshalIndent(state, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(s.dataFile, data, 0644)
}

func (s *Store) GetState(ctx context.Context) (*domain.AppState, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	itemsCopy := make([]domain.Item, len(s.items))
	copy(itemsCopy, s.items)

	wishesCopy := make([]domain.Wish, len(s.wishes))
	copy(wishesCopy, s.wishes)

	questionsCopy := make([]domain.Question, len(s.questions))
	copy(questionsCopy, s.questions)

	recognitionsCopy := make([]domain.Recognition, len(s.recognitions))
	copy(recognitionsCopy, s.recognitions)

	categoriesCopy := make([]domain.Category, len(s.categories))
	copy(categoriesCopy, s.categories)

	return &domain.AppState{
		Categories:   categoriesCopy,
		Items:        itemsCopy,
		Wishes:       wishesCopy,
		Questions:    questionsCopy,
		Recognitions: recognitionsCopy,
	}, nil
}

func (s *Store) GetItems(ctx context.Context) ([]domain.Item, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	itemsCopy := make([]domain.Item, len(s.items))
	copy(itemsCopy, s.items)
	return itemsCopy, nil
}

func (s *Store) GetItemByID(ctx context.Context, id string) (*domain.Item, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	for _, it := range s.items {
		if it.ID == id {
			c := it
			return &c, nil
		}
	}
	return nil, errors.New("item not found")
}

func (s *Store) SaveItem(ctx context.Context, item domain.Item) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	found := false
	for i, it := range s.items {
		if it.ID == item.ID {
			s.items[i] = item
			found = true
			break
		}
	}
	if !found {
		s.items = append([]domain.Item{item}, s.items...)
	}

	return s.persist()
}

func (s *Store) ToggleItemLike(ctx context.Context, itemID, userID string) (*domain.Item, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	for i := range s.items {
		if s.items[i].ID == itemID {
			likes := s.items[i].Likes
			liked := false
			newLikes := make([]string, 0, len(likes))
			for _, u := range likes {
				if u == userID {
					liked = true
				} else {
					newLikes = append(newLikes, u)
				}
			}
			if !liked {
				newLikes = append(newLikes, userID)
			}
			s.items[i].Likes = newLikes
			_ = s.persist()
			result := s.items[i]
			return &result, nil
		}
	}
	return nil, errors.New("item not found")
}

func (s *Store) GetWishes(ctx context.Context) ([]domain.Wish, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	wishesCopy := make([]domain.Wish, len(s.wishes))
	copy(wishesCopy, s.wishes)
	return wishesCopy, nil
}

func (s *Store) SaveWish(ctx context.Context, wish domain.Wish) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	found := false
	for i, w := range s.wishes {
		if w.ID == wish.ID {
			s.wishes[i] = wish
			found = true
			break
		}
	}
	if !found {
		s.wishes = append([]domain.Wish{wish}, s.wishes...)
	}

	return s.persist()
}

func (s *Store) GetQuestions(ctx context.Context) ([]domain.Question, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	qCopy := make([]domain.Question, len(s.questions))
	copy(qCopy, s.questions)
	return qCopy, nil
}

func (s *Store) SaveQuestion(ctx context.Context, question domain.Question) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	found := false
	for i, q := range s.questions {
		if q.ID == question.ID {
			s.questions[i] = question
			found = true
			break
		}
	}
	if !found {
		s.questions = append([]domain.Question{question}, s.questions...)
	}

	return s.persist()
}

func (s *Store) AddReply(ctx context.Context, questionID string, reply domain.Reply) (*domain.Question, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	for i := range s.questions {
		if s.questions[i].ID == questionID {
			s.questions[i].Replies = append(s.questions[i].Replies, reply)
			_ = s.persist()
			result := s.questions[i]
			return &result, nil
		}
	}
	return nil, errors.New("question not found")
}

func (s *Store) ToggleHelpful(ctx context.Context, questionID, replyID, authorName string) (*domain.Question, *domain.Recognition, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	for i := range s.questions {
		if s.questions[i].ID == questionID {
			var targetReply *domain.Reply
			for j := range s.questions[i].Replies {
				if s.questions[i].Replies[j].ID == replyID {
					targetReply = &s.questions[i].Replies[j]
					targetReply.HelpfulByAuthor = !targetReply.HelpfulByAuthor
					break
				}
			}

			if targetReply == nil {
				return nil, nil, errors.New("reply not found")
			}

			var affectedRecognition *domain.Recognition
			if targetReply.HelpfulByAuthor {
				rec := domain.Recognition{
					RewardID:             fmt.Sprintf("rec_%d", time.Now().UnixNano()),
					QuestionID:           questionID,
					QuestionTitle:        s.questions[i].Title,
					RecipientID:          targetReply.AuthorID,
					RecipientName:        targetReply.AuthorName,
					MarkedByAuthorName:   authorName,
					SupportingReplyIDs:   []string{replyID},
					SupportingVersionIDs: []string{targetReply.VersionID},
					Status:               "active_public",
					Archived:             false,
					ConsentVersion:       "v1",
					UpdatedAt:            time.Now().Format("2006-01-02 15:04"),
				}
				s.recognitions = append([]domain.Recognition{rec}, s.recognitions...)
				affectedRecognition = &rec
			} else {
				for k := range s.recognitions {
					if s.recognitions[k].QuestionID == questionID && s.recognitions[k].RecipientID == targetReply.AuthorID {
						s.recognitions[k].Status = "inactive"
						affectedRecognition = &s.recognitions[k]
						break
					}
				}
			}

			_ = s.persist()
			qResult := s.questions[i]
			return &qResult, affectedRecognition, nil
		}
	}

	return nil, nil, errors.New("question not found")
}
