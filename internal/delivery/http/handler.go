package http

import (
	"encoding/json"
	"net/http"

	"retro/internal/domain"
	"retro/internal/usecase"
)

type Handler struct {
	uc usecase.CatalogUseCase
}

func NewHandler(uc usecase.CatalogUseCase) *Handler {
	return &Handler{uc: uc}
}

func (h *Handler) writeJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(data)
}

func (h *Handler) writeError(w http.ResponseWriter, status int, msg string) {
	h.writeJSON(w, status, map[string]string{"error": msg})
}

func (h *Handler) GetState(w http.ResponseWriter, r *http.Request) {
	state, err := h.uc.GetState(r.Context())
	if err != nil {
		h.writeError(w, http.StatusInternalServerError, err.Error())
		return
	}
	h.writeJSON(w, http.StatusOK, state)
}

func (h *Handler) GetItems(w http.ResponseWriter, r *http.Request) {
	items, err := h.uc.GetItems(r.Context())
	if err != nil {
		h.writeError(w, http.StatusInternalServerError, err.Error())
		return
	}
	h.writeJSON(w, http.StatusOK, items)
}

func (h *Handler) CreateItem(w http.ResponseWriter, r *http.Request) {
	var item domain.Item
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		h.writeError(w, http.StatusBadRequest, "invalid request body")
		return
	}
	created, err := h.uc.CreateItem(r.Context(), item)
	if err != nil {
		h.writeError(w, http.StatusBadRequest, err.Error())
		return
	}
	h.writeJSON(w, http.StatusCreated, created)
}

type LikeRequest struct {
	ItemID string `json:"itemId"`
	UserID string `json:"userId"`
}

func (h *Handler) ToggleLike(w http.ResponseWriter, r *http.Request) {
	var req LikeRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.writeError(w, http.StatusBadRequest, "invalid request body")
		return
	}
	item, err := h.uc.ToggleLike(r.Context(), req.ItemID, req.UserID)
	if err != nil {
		h.writeError(w, http.StatusBadRequest, err.Error())
		return
	}
	h.writeJSON(w, http.StatusOK, item)
}

func (h *Handler) CreateWish(w http.ResponseWriter, r *http.Request) {
	var wish domain.Wish
	if err := json.NewDecoder(r.Body).Decode(&wish); err != nil {
		h.writeError(w, http.StatusBadRequest, "invalid request body")
		return
	}
	created, err := h.uc.CreateWish(r.Context(), wish)
	if err != nil {
		h.writeError(w, http.StatusBadRequest, err.Error())
		return
	}
	h.writeJSON(w, http.StatusCreated, created)
}

func (h *Handler) CreateQuestion(w http.ResponseWriter, r *http.Request) {
	var q domain.Question
	if err := json.NewDecoder(r.Body).Decode(&q); err != nil {
		h.writeError(w, http.StatusBadRequest, "invalid request body")
		return
	}
	created, err := h.uc.CreateQuestion(r.Context(), q)
	if err != nil {
		h.writeError(w, http.StatusBadRequest, err.Error())
		return
	}
	h.writeJSON(w, http.StatusCreated, created)
}

type ReplyRequest struct {
	QuestionID string       `json:"questionId"`
	Reply      domain.Reply `json:"reply"`
}

func (h *Handler) AddReply(w http.ResponseWriter, r *http.Request) {
	var req ReplyRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.writeError(w, http.StatusBadRequest, "invalid request body")
		return
	}
	q, err := h.uc.AddReply(r.Context(), req.QuestionID, req.Reply)
	if err != nil {
		h.writeError(w, http.StatusBadRequest, err.Error())
		return
	}
	h.writeJSON(w, http.StatusOK, q)
}

type HelpfulRequest struct {
	QuestionID string `json:"questionId"`
	ReplyID    string `json:"replyId"`
	AuthorName string `json:"authorName"`
}

func (h *Handler) ToggleHelpful(w http.ResponseWriter, r *http.Request) {
	var req HelpfulRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.writeError(w, http.StatusBadRequest, "invalid request body")
		return
	}
	q, rec, err := h.uc.ToggleHelpful(r.Context(), req.QuestionID, req.ReplyID, req.AuthorName)
	if err != nil {
		h.writeError(w, http.StatusBadRequest, err.Error())
		return
	}
	h.writeJSON(w, http.StatusOK, map[string]interface{}{
		"question":    q,
		"recognition": rec,
	})
}
