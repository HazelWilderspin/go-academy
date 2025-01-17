package crud

import (
	"sync"

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

type Store struct {
	Data []User
	Lock sync.Mutex
}
