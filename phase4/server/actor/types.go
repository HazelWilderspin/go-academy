package actor

import (
	"server/crud"

	"github.com/google/uuid"
)

type Response struct {
	GetUserResponse crud.User
}

type Request struct {
	ResponseChannel chan Response
	ErrorChannel    chan error
	Action          string
	GetUserArg      string
	PostListArgs
	DeleteListArgs
	PostItemArgs
	DeleteItemArgs
}

type PostListArgs struct {
	userId  uuid.UUID
	newList crud.List
}

type DeleteListArgs struct {
	userId uuid.UUID
	listId uuid.UUID
}

type PostItemArgs struct {
	userId  uuid.UUID
	listId  uuid.UUID
	newItem crud.Item
}

type DeleteItemArgs struct {
	userId uuid.UUID
	listId uuid.UUID
	itemId uuid.UUID
}
