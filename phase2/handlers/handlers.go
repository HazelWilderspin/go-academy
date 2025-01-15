package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	d "phase2/data"
	t "phase2/types"
	"time"
)

var UNMARSHALED_DATA_STORE t.User
var USERNAME string

type Body struct {
	Username string
}

func LoadData(writer http.ResponseWriter, req *http.Request) {
	fmt.Println("##### Load data handler begun #####")
	defer fmt.Println("##### Load data handler ended #####")

	var unmarshaledBody Body

	byteArr, err := io.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(byteArr, &unmarshaledBody)
	if err != nil {
		slog.Error("Failed to unmarshal json", "err", err)
	}

	USERNAME = unmarshaledBody.Username

	ctx := req.Context()
	// subCtx, cancel := context.WithCancel(ctx)

	select {
	case <-time.After(1 * time.Second):
		// fmt.Println("Testing 1 sec channel info")

		data, err := d.LoadMockData(USERNAME)
		if err != nil {
			slog.Error("Data failed to initialize", "Server Error: ", err)
		}

		UNMARSHALED_DATA_STORE = data

		writer.Write([]byte(UNMARSHALED_DATA_STORE.Forename))

	case <-ctx.Done():
		err := ctx.Err()
		slog.Info("Context done called", "Server Error: ", err)

		internalError := http.StatusInternalServerError
		http.Error(writer, err.Error(), internalError)
	}
}

func ListHeaders(writer http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for _, header := range headers {
			fmt.Fprintf(writer, "%v: %v\n", name, header)
		}
	}
}

func GetUser(writer http.ResponseWriter, req *http.Request) {
	fmt.Println("##### Get user data handler begun #####")
	defer fmt.Println("##### Get user data handler ended #####")

	byteArr, err := json.MarshalIndent(&UNMARSHALED_DATA_STORE, "  ", "  ")
	if err != nil {
		slog.Error("Failed to marshal json", "err", err)
	}

	writer.Write(byteArr)
}

// func GetLists(writer http.ResponseWriter, req *http.Request) {
// 	slog.Info("##### User lists get request #####")

// 	fmt.Println(UNMARSHALED_DATA_STORE)

// 	byteArr, err := json.MarshalIndent(&UNMARSHALED_DATA_STORE, "  ", "  ")
// 	if err != nil {
// 		slog.Error("Failed to marshal json", "err", err)
// 	}

// 	writer.Write(byteArr)
// }

// func GetList(writer http.ResponseWriter, req *http.Request) {
// 	slog.Info("##### User lists get request #####")

// 	fmt.Println(UNMARSHALED_DATA_STORE)

// 	byteArr, err := json.MarshalIndent(&UNMARSHALED_DATA_STORE, "  ", "  ")
// 	if err != nil {
// 		slog.Error("Failed to marshal json", "err", err)
// 	}

// 	writer.Write(byteArr)
// }
