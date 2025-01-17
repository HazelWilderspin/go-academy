package handlers

import (
	"context"
	"log/slog"
	t "phase2/crud"

	"github.com/google/uuid"
)

type GetUserRequestBody struct {
	UserId uuid.UUID `json:"user_detail_id"`
}

type GetListRequestBody struct {
	UserId uuid.UUID `json:"user_detail_id"`
	ListId uuid.UUID `json:"list_id"`
}

type PostListRequestBody struct {
	UserId  uuid.UUID `json:"user_detail_id"`
	NewList t.List    `json:"list"`
}

type PutListNameRequestBody struct {
	UserId      uuid.UUID `json:"user_detail_id"`
	ListId      uuid.UUID `json:"list_id"`
	NewListName string    `json:"new_list_name"`
}

type PutListCompletionRequestBody struct {
	UserId         uuid.UUID `json:"user_detail_id"`
	ListId         uuid.UUID `json:"list_id"`
	ListIsComplete bool      `json:"list_is_complete"`
}

type DeleteListRequestBody struct {
	UserId uuid.UUID `json:"user_detail_id"`
	ListId uuid.UUID `json:"list_id"`
}

type GetItemRequestBody struct {
	UserId uuid.UUID `json:"user_detail_id"`
	ListId uuid.UUID `json:"list_id"`
	ItemId uuid.UUID `json:"item_id"`
}

type PostItemRequestBody struct {
	UserId  uuid.UUID `json:"user_detail_id"`
	ListId  uuid.UUID `json:"list_id"`
	NewItem t.Item    `json:"item"`
}

type PutItemRequestBody struct {
	UserId uuid.UUID `json:"user_detail_id"`
	ListId uuid.UUID `json:"list_id"`
	Item   t.Item    `json:"item"`
}

type DeleteItemRequestBody struct {
	UserId uuid.UUID `json:"user_detail_id"`
	ListId uuid.UUID `json:"list_id"`
	ItemId uuid.UUID `json:"item_id"`
}

type ContextHandler struct {
	slog.Handler
}

func (h *ContextHandler) ContextHandlerReceiver(ctx context.Context, r slog.Record) error {
	traceId, ok := ctx.Value("trace_id").(string)

	if ok {
		r.AddAttrs(slog.String("trace_id", traceId))
	}

	return h.Handler.Handle(ctx, r)
}

type CtxKey string
