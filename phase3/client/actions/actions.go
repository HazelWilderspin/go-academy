package actions

import (
	"bufio"
	"bytes"
	"errors"
	"log/slog"

	"fmt"
	"net/http"
)

var clt http.Client

func GetUser(reqBody []byte) ([]byte, error) {
	fmt.Println("GetUser action called")

	resp, err := clt.Post("http://localhost:8080/GetUser", "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		slog.Error(err.Error())
		return nil, err
	}

	defer resp.Body.Close()

	var marshalledData []byte

	scanner := bufio.NewScanner(resp.Body)
	for scanner.Scan() {
		marshalledData = append(marshalledData, scanner.Bytes()...)
	}
	if scanner.Err() != nil {
		return nil, scanner.Err()
	}
	if resp.StatusCode != 200 {
		err = errors.New(resp.Status)
	}

	return marshalledData, err
}

func PostList(reqBody []byte) error {
	fmt.Println("PostList action called")

	resp, err := clt.Post("http://localhost:8080/PostList", "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		err = errors.New(resp.Status)
	}

	return err
}

func DeleteList(reqBody []byte) error {
	fmt.Println("DeleteList action called")

	resp, err := clt.Post("http://localhost:8080/DeleteList", "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		err = errors.New(resp.Status)
	}

	return err
}

func PostItem(reqBody []byte) error {
	fmt.Println("PostItem action called")

	resp, err := clt.Post("http://localhost:8080/PostItem", "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		err = errors.New(resp.Status)
	}

	return err
}

func DeleteItem(reqBody []byte) error {
	fmt.Println("DeleteItem action called")

	resp, err := clt.Post("http://localhost:8080/DeleteItem", "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		err = errors.New(resp.Status)
	}

	return err
}
