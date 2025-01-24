package actor

import (
	"errors"
	"fmt"
	"server/crud"

	"github.com/google/uuid"
)

var RequestChannel = make(chan Request, 100)

func Actor() {
	fmt.Println("Starting up RequestChannel listener")

	for request := range RequestChannel {
		fmt.Println("--- Actor Received action from Request channel: ", request.Action)

		switch request.Action {
		case "GetUser":
			user, err := crud.ReadUser(request.GetUserArg)
			request.ResponseChannel <- Response{user, err}

		case "PostList":
			err := crud.CreateList(request.PostListArgs.userId, request.PostListArgs.newList)
			request.ResponseChannel <- Response{crud.User{}, err}

		case "DeleteList":
			err := crud.DeleteList(request.DeleteListArgs.userId, request.DeleteListArgs.listId)
			request.ResponseChannel <- Response{crud.User{}, err}

		case "PostItem":
			err := crud.CreateItem(request.PostItemArgs.userId, request.PostItemArgs.listId, request.PostItemArgs.newItem)
			request.ResponseChannel <- Response{crud.User{}, err}

		case "PutItem":
			_, err := crud.UpdateItem(request.PostItemArgs.userId, request.PostItemArgs.listId, request.PostItemArgs.newItem)
			request.ResponseChannel <- Response{crud.User{}, err}

		case "DeleteItem":
			err := crud.DeleteItem(request.DeleteItemArgs.userId, request.DeleteItemArgs.listId, request.DeleteItemArgs.itemId)
			request.ResponseChannel <- Response{crud.User{}, err}

		default:
			request.ResponseChannel <- Response{crud.User{}, errors.New("Actor defaulted, request action not viable")}
		}
	}
}

func AddGetUserToRequestChannel(username string, action string) (crud.User, error) {

	request := Request{
		ResponseChannel: make(chan Response, 1),
		Action:          action,
		GetUserArg:      username,
	}

	RequestChannel <- request
	response := <-request.ResponseChannel
	return response.GetUserResponse, response.Err
}

func AddPostListToRequestChannel(userId uuid.UUID, newList crud.List, action string) error {

	request := Request{
		ResponseChannel: make(chan Response, 1),
		Action:          action,
		PostListArgs:    PostListArgs{userId, newList},
	}

	RequestChannel <- request
	response := <-request.ResponseChannel
	return response.Err
}

func AddDeleteListToRequestChannel(userId uuid.UUID, listId uuid.UUID, action string) error {

	request := Request{
		ResponseChannel: make(chan Response, 1),
		Action:          action,
		DeleteListArgs:  DeleteListArgs{userId, listId},
	}

	RequestChannel <- request
	response := <-request.ResponseChannel
	return response.Err
}

func AddPostItemToRequestChannel(userId uuid.UUID, listId uuid.UUID, newItem crud.Item, action string) error {

	request := Request{
		ResponseChannel: make(chan Response, 1),
		Action:          action,
		PostItemArgs:    PostItemArgs{userId, listId, newItem},
	}

	RequestChannel <- request
	response := <-request.ResponseChannel
	return response.Err
}

func AddPutItemToRequestChannel(userId uuid.UUID, listId uuid.UUID, newItem crud.Item, action string) error {

	request := Request{
		ResponseChannel: make(chan Response, 1),

		Action:       action,
		PostItemArgs: PostItemArgs{userId, listId, newItem},
	}

	RequestChannel <- request
	response := <-request.ResponseChannel
	return response.Err
}

func AddDeleteItemToRequestChannel(userId uuid.UUID, listId uuid.UUID, itemId uuid.UUID, action string) error {

	request := Request{
		ResponseChannel: make(chan Response, 1),

		Action:         action,
		DeleteItemArgs: DeleteItemArgs{userId, listId, itemId},
	}

	RequestChannel <- request
	response := <-request.ResponseChannel
	return response.Err
}
