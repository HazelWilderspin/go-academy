package handlers

import (
	"github.com/google/uuid"
)

type Item struct {
	ItemId        uuid.UUID `json:"item_id"`
	ItemName      string    `json:"item_name"`
	ItemDesc      string    `json:"item_description"`
	ItemIsChecked bool      `json:"item_is_checked"`
}

type List struct {
	ListId     uuid.UUID `json:"list_id"`
	ListName   string    `json:"list_name"`
	InitDate   string    `json:"list_init_date"`
	IsComplete bool      `json:"list_is_complete"`
	Items      []Item    `json:"items"`
}

type User struct {
	UserDetailId    uuid.UUID `json:"user_detail_id"`
	UserName        string    `json:"username"`
	Forename        string    `json:"forename"`
	UserPermissions string    `json:"user_permissions"`
	InitDate        string    `json:"init_date"`
	Lists           []List    `json:"lists"`
}

type Cache struct {
	UserDetailId uuid.UUID
	Username     string
	Forename     string
	ListCount    int
	Lists        []List
}

type GetUserRequestBody struct {
	Username string `json:"username"`
}

type GetUserResponseBody struct {
	UserDetailId    uuid.UUID `json:"user_detail_id"`
	UserName        string    `json:"username"`
	Forename        string    `json:"forename"`
	UserPermissions string    `json:"user_permissions"`
	InitDate        string    `json:"init_date"`
	Lists           []List    `json:"lists"`
}

type PostListRequestBody struct {
	UserDetailId uuid.UUID `json:"user_detail_id"`
	NewList      List      `json:"list"`
}

type DeleteListRequestBody struct {
	UserDetailId uuid.UUID `json:"user_detail_id"`
	ListId       uuid.UUID `json:"list_id"`
}

type AddItemRequestBody struct {
	UserDetailId uuid.UUID `json:"user_detail_id"`
	ListId       uuid.UUID `json:"list_id"`
	NewItem      Item      `json:"item"`
}

type DeleteItemRequestBody struct {
	UserDetailId uuid.UUID `json:"user_detail_id"`
	ListId       uuid.UUID `json:"list_id"`
	ItemId       uuid.UUID `json:"item_id"`
}
