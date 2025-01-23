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
			if err != nil {
				request.ErrorChannel <- err
			}

			request.ResponseChannel <- Response{user}

		case "PostList":

			err := crud.CreateList(request.PostListArgs.userId, request.PostListArgs.newList)
			if err != nil {
				request.ErrorChannel <- err
			}

		case "DeleteList":

			err := crud.DeleteList(request.DeleteListArgs.userId, request.DeleteListArgs.listId)
			if err != nil {
				request.ErrorChannel <- err
			}

		case "PostItem":

			err := crud.CreateItem(request.PostItemArgs.userId, request.PostItemArgs.listId, request.PostItemArgs.newItem)
			if err != nil {
				request.ErrorChannel <- err
			}

		case "PutItem":

			_, err := crud.UpdateItem(request.PostItemArgs.userId, request.PostItemArgs.listId, request.PostItemArgs.newItem)
			if err != nil {
				request.ErrorChannel <- err
			}

		case "DeleteItem":

			err := crud.DeleteItem(request.DeleteItemArgs.userId, request.DeleteItemArgs.listId, request.DeleteItemArgs.itemId)
			if err != nil {
				request.ErrorChannel <- err
			}

		default:
			request.ErrorChannel <- errors.New("Actor defaulted, request action not viable")
		}
		close(request.ErrorChannel)
		close(request.ResponseChannel)
	}
}

func AddGetUserToRequestChannel(username string, action string) (crud.User, error) {

	request := Request{
		ResponseChannel: make(chan Response),
		ErrorChannel:    make(chan error),
		Action:          action,
		GetUserArg:      username,
	}

	RequestChannel <- request

	select {
	case err := <-request.ErrorChannel:
		return crud.User{}, err
	case response := <-request.ResponseChannel:
		return response.GetUserResponse, nil
	}

}

func AddPostListToRequestChannel(userId uuid.UUID, newList crud.List, action string) error {

	request := Request{
		ResponseChannel: make(chan Response),
		ErrorChannel:    make(chan error),
		Action:          action,
		PostListArgs:    PostListArgs{userId, newList},
	}

	RequestChannel <- request
	return <-request.ErrorChannel

}

func AddDeleteListToRequestChannel(userId uuid.UUID, listId uuid.UUID, action string) error {

	request := Request{
		ResponseChannel: make(chan Response),
		ErrorChannel:    make(chan error),
		Action:          action,
		DeleteListArgs:  DeleteListArgs{userId, listId},
	}

	RequestChannel <- request
	return <-request.ErrorChannel

}

func AddPostItemToRequestChannel(userId uuid.UUID, listId uuid.UUID, newItem crud.Item, action string) error {

	request := Request{
		ResponseChannel: make(chan Response),
		ErrorChannel:    make(chan error),
		Action:          action,
		PostItemArgs:    PostItemArgs{userId, listId, newItem},
	}

	RequestChannel <- request
	return <-request.ErrorChannel

}

func AddPutItemToRequestChannel(userId uuid.UUID, listId uuid.UUID, newItem crud.Item, action string) error {

	request := Request{
		ResponseChannel: make(chan Response),
		ErrorChannel:    make(chan error),
		Action:          action,
		PostItemArgs:    PostItemArgs{userId, listId, newItem},
	}

	RequestChannel <- request
	return <-request.ErrorChannel

}

func AddDeleteItemToRequestChannel(userId uuid.UUID, listId uuid.UUID, itemId uuid.UUID, action string) error {

	request := Request{
		ResponseChannel: make(chan Response),
		ErrorChannel:    make(chan error),
		Action:          action,
		DeleteItemArgs:  DeleteItemArgs{userId, listId, itemId},
	}

	RequestChannel <- request
	return <-request.ErrorChannel

}
