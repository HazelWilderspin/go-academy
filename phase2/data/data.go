package data

import (
	"bufio"
	"encoding/json"
	"log/slog"
	"os"
	t "phase2/types"
)

const FILE_PATH string = "data/mockData.json"

func LoadMockData(username string) (t.User, error) {
	var marshalledData []byte
	var unmarshaledData []t.User
	var userData t.User

	file, err := os.Open(FILE_PATH)
	if err != nil {
		slog.Error("Opening file was unsuccessful", "Server Error: ", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		marshalledData = append(marshalledData, scanner.Bytes()...)
	}
	if err != nil {
		slog.Error("bufio scanner returned an error", "Server Error: ", scanner.Err())
	}

	err = json.Unmarshal(marshalledData, &unmarshaledData)
	if err != nil {
		slog.Error("Failed to unmarshal json", "Server Error: ", err)
	}

	for _, user := range unmarshaledData {
		if user.UserName == username {
			userData = user
		}
	}

	return userData, err
}
