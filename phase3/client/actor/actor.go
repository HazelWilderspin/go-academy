package actor

import (
	"client/actions"
	"fmt"
	"log/slog"
)

var RequestChannel = make(chan Request, 100)

func Actor() {
	fmt.Println("Starting up RequestChannel listener")

	for request := range RequestChannel {
		fmt.Println("--- Actor Received action from Request channel: ", request.Action)

		switch request.Action {
		case "GetUser":

			byteArr, err := actions.GetUser(request.Args)
			request.ResponseChannel <- Response{byteArr, err}

		case "PostList":

			err := actions.PostList(request.Args)
			request.ResponseChannel <- Response{nil, err}

		case "DeleteList":

			err := actions.DeleteList(request.Args)
			request.ResponseChannel <- Response{nil, err}

		case "PostItem":

			err := actions.PostItem(request.Args)
			request.ResponseChannel <- Response{nil, err}

		case "DeleteItem":

			err := actions.DeleteItem(request.Args)
			request.ResponseChannel <- Response{nil, err}

		default:
			slog.Error("Actor defaulted, request action not viable")
		}
	}
}

func AddRequestToRequestChannel(reqBody []byte, action string) ([]byte, error) {

	request := Request{
		ResponseChannel: make(chan Response),
		Action:          action,
		Args:            reqBody,
	}

	// timer/timeout and return timeout error ?

	fmt.Println("Adding request to RequestChannel")
	RequestChannel <- request

	fmt.Println("Pulling marshalled data from ResponseChannel")
	response := <-request.ResponseChannel

	fmt.Println("--- API response received!")
	close(request.ResponseChannel)

	return response.Data, response.Err
}
