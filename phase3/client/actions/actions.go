package rest

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

func GetUser(username string) (User, error) {
	var user User
	var marshalledData []byte

	reqBody, err := json.Marshal(GetUserRequestBody{username})
	if err != nil {
		return User{}, err
	}

	resp, err := http.Post("http://localhost:8080/GetUser", "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		return User{}, err
	}
	defer resp.Body.Close()

	fmt.Println("Response status:", resp.Status)

	scanner := bufio.NewScanner(resp.Body)
	for scanner.Scan() {
		marshalledData = append(marshalledData, scanner.Bytes()...)
	}
	if err := scanner.Err(); err != nil {
		return User{}, err
	}

	err = json.Unmarshal(marshalledData, &user)
	if err != nil {
		return User{}, err
	}

	return user, err
}

func PostList(UserDetailId *uuid.UUID, NewList List) (int, error) {
	fmt.Printf("UserId: %v\n NewList: %v\n", UserDetailId, NewList)
	reqBody, err := json.Marshal(PostListRequestBody{*UserDetailId, NewList})
	if err != nil {
		return 0, err
	}

	resp, err := http.Post("http://localhost:8080/PostList", "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	fmt.Println("Response status:", resp.Status)

	return resp.StatusCode, err
}
