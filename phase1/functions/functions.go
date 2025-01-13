package functions

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	t "phase1/types"

	"github.com/google/uuid"
)

func ReportIfError(msg string, err error) {
	if err != nil {
		log.Fatal(err.Error())
		fmt.Println(msg)
	}
}

func Turn[T any](condition bool, trueValue, falseValue T) T {
	if condition {
		return trueValue
	}
	return falseValue
}

func CreateItem() []t.Item {
	FILE_PATH := "data2.json"

	data := GetData(FILE_PATH)

	var items []t.Item
	err := json.Unmarshal(data, &items)
	ReportIfError("Failed to unmarshal json", err)
	fmt.Println("Number of items retrieved from data2: ", len(items))

	newItem := t.Item{uuid.New(), "New item", "New item description", false}

	items = append(items, newItem)

	byteArr, err := json.Marshal(items)
	ReportIfError("Failed to marshal json", err)

	// options := os.O_WRONLY | os.O_APPEND | os.O_CREATE
	// file, err := os.OpenFile(FILE_PATH, options, os.FileMode(0600))
	// ReportIfError(err)
	// defer file.Close()

	err = os.WriteFile(FILE_PATH, byteArr, 0644)
	ReportIfError("Writing your byte array to the file selected failed", err)
	return items
}

func CreateList() []t.List {
	FILE_PATH := "data3.json"

	data := GetListData(FILE_PATH, "New list")

	var lists []t.List
	err := json.Unmarshal(data, &lists)
	ReportIfError("Failed to unmarshal json", err)
	fmt.Println("Number of lists retrieved from data3: ", len(lists))

	var items []t.Item
	newItem := t.Item{uuid.New(), "New item", "New item description", false}
	items = append(items, newItem)

	newList := t.List{uuid.New(), "New list", "2012-04-23T18:25:43.511Z", false, items}
	lists = append(lists, newList)

	// Use a different method to format the json and pass in a diff arg
	byteArr, err := json.Marshal(lists)
	ReportIfError("Failed to marshal json", err)

	err = os.WriteFile(FILE_PATH, byteArr, 0644)
	ReportIfError("Writing your byte array to the file selected failed", err)
	return lists
}

func CreateUser(username string) []t.User {
	FILE_PATH := "data4.json"

	data := GetUserData(FILE_PATH)

	var users []t.User
	err := json.Unmarshal(data, &users)
	ReportIfError("Failed to unmarshal json", err)
	fmt.Println("Number of users retrieved from data4: ", len(users))

	var items []t.Item
	newItem := t.Item{uuid.New(), "New item", "New item description", false}
	items = append(items, newItem)

	var lists []t.List
	newList := t.List{uuid.New(), "New list", "2012-04-23T18:25:43.511Z", false, items}
	lists = append(lists, newList)

	newUser := t.User{uuid.New(), username, "Alice", "CREATE_READ_UPDATE_DELETE", "2012-04-23T18:25:43.511Z", lists}
	users = append(users, newUser)

	byteArr, err := json.Marshal(users)
	ReportIfError("Failed to marshal json", err)

	// options := os.O_WRONLY | os.O_APPEND | os.O_CREATE
	// file, err := os.OpenFile(FILE_PATH, options, os.FileMode(0600))
	// ReportIfError(err)
	// defer file.Close()

	err = os.WriteFile(FILE_PATH, byteArr, 0644)
	ReportIfError("Writing your byte array to the file selected failed", err)
	return users
}

func GetData(fileName string) []byte {
	var data []byte

	file, err := os.Open(fileName)
	ReportIfError("Opening file was unsuccessful", err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data = append(data, scanner.Bytes()...)
	}
	ReportIfError("bufio scanner returned an error", scanner.Err())

	return data
}

func GetListData(fileName string, listName string) []byte {
	var data []byte

	file, err := os.Open(fileName)
	ReportIfError("Opening file was unsuccessful", err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data = append(data, scanner.Bytes()...)
	}
	ReportIfError("bufio scanner returned an error", scanner.Err())

	// filter data

	return data
}

func GetUserData(fileName string) []byte {

	var data []byte

	file, err := os.Open(fileName)
	ReportIfError("Opening file was unsuccessful", err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data = append(data, scanner.Bytes()...)
	}
	ReportIfError("bufio scanner returned an error", scanner.Err())

	// filter data

	return data
}
