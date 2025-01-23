package handlers

import (
	"client/actions"

	"github.com/google/uuid"
)

type Cache struct {
	UserDetailId uuid.UUID
	Username     string
	Forename     string
	ListCount    int
	Lists        []actions.List
}

type GetUserRequestBody struct {
	Username string `json:"username"`
}

type GetUserResponseBody struct {
	UserDetailId    uuid.UUID      `json:"user_detail_id"`
	UserName        string         `json:"username"`
	Forename        string         `json:"forename"`
	UserPermissions string         `json:"user_permissions"`
	InitDate        string         `json:"init_date"`
	Lists           []actions.List `json:"lists"`
}

type PostListRequestBody struct {
	UserDetailId uuid.UUID    `json:"user_detail_id"`
	NewList      actions.List `json:"list"`
}

type DeleteListRequestBody struct {
	UserDetailId uuid.UUID `json:"user_detail_id"`
	ListId       uuid.UUID `json:"list_id"`
}

type AddItemRequestBody struct {
	UserDetailId uuid.UUID    `json:"user_detail_id"`
	ListId       uuid.UUID    `json:"list_id"`
	NewItem      actions.Item `json:"item"`
}

type DeleteItemRequestBody struct {
	UserDetailId uuid.UUID `json:"user_detail_id"`
	ListId       uuid.UUID `json:"list_id"`
	ItemId       uuid.UUID `json:"item_id"`
}
