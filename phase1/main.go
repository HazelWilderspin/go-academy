package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	f "phase1/functions"
	r "phase1/repl"
	t "phase1/types"
	"strings"
)

const FILE_PATH string = "data4.json"
const USERNAME string = "ALPHA"

var UNMARSHALED_DATA []t.User

// var USERNAME string

// receivers?
// func (u User) readLists() []List {
// 	return u.Lists
// }

func main() {
	reader := bufio.NewReader(os.Stdin)

	// fmt.Println("Please enter your username:")
	// username, err := reader.ReadString('\n')
	// f.ReportIfError("bufio input resulted in an error", err)
	// USERNAME = strings.TrimSpace(USERNAME)

	name, lists := reloadData(FILE_PATH, USERNAME)
	fmt.Printf("Hello %s \n You have %d list(s), their names are as follows: \n", name, len(lists))

	for _, value := range lists {
		fmt.Printf("    %v\n", value.ListName)
	}

	repeatListOptions(reader)
}

func listActions(command string, reader *bufio.Reader) {
	listsReadOnly := GetLists()

	switch strings.TrimSpace(command) {
	case "createList":
		newList := r.FormNewList()

		for key, user := range UNMARSHALED_DATA {
			if user.UserName == USERNAME {
				UNMARSHALED_DATA[key].Lists = append(UNMARSHALED_DATA[key].Lists, newList)
			}
		}

		SaveFile()
		reloadData(FILE_PATH, USERNAME)
		repeatListOptions(reader)

	case "readLists":
		fmt.Printf("You have %d list(s), their names are as follows: \n", len(listsReadOnly))

		for _, value := range listsReadOnly {
			fmt.Printf("    %v\n", value.ListName)
		}
		repeatListOptions(reader)

	case "readList":
		fmt.Println("Enter list name to read:")
		listName, err := reader.ReadString('\n')
		f.ReportIfError("bufio input resulted in an error", err)

		for _, list := range listsReadOnly {
			if list.ListName == strings.TrimSpace(listName) {
				items := list.Items
				for _, item := range items {
					complete := f.Turn(item.ItemIsChecked, "Complete", "Incomplete")
					fmt.Printf("    %v : %v : %v\n", item.ItemName, item.ItemDesc, complete)
				}
			}
		}

		repeatListOptions(reader)

	case "deleteList":
		fmt.Printf("You have %d list(s), their names are as follows: \n", len(listsReadOnly))
		for _, value := range listsReadOnly {
			fmt.Printf("    %v\n", value.ListName)
		}

		fmt.Println("Enter list name to delete it:")
		listName, err := reader.ReadString('\n')
		f.ReportIfError("bufio input resulted in an error", err)

		for key1, user := range UNMARSHALED_DATA {
			if user.UserName == USERNAME {
				for key2, list := range user.Lists {
					if list.ListName == strings.TrimSpace(listName) {
						if key2 >= len(user.Lists) {
							UNMARSHALED_DATA[key1].Lists = UNMARSHALED_DATA[key1].Lists[:key2]
						} else {
							UNMARSHALED_DATA[key1].Lists = append(UNMARSHALED_DATA[key1].Lists[:key2], UNMARSHALED_DATA[key1].Lists[key2+1:]...)
						}
					}
				}
			}
		}

		SaveFile()
		reloadData(FILE_PATH, USERNAME)
		repeatListOptions(reader)

	case "exit":
		fmt.Println("Goodbye")
	}
}

func repeatListOptions(reader *bufio.Reader) {
	r.PrintListOptions()

	command, err := reader.ReadString('\n')
	f.ReportIfError("bufio input resulted in an error", err)
	listActions(strings.TrimSpace(command), reader)
}

func reloadData(fileName string, username string) (string, []t.List) {
	var marshalledData []byte
	var target t.User

	file, err := os.Open(fileName)
	f.ReportIfError("Opening file was unsuccessful", err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		marshalledData = append(marshalledData, scanner.Bytes()...)
	}
	f.ReportIfError("bufio scanner returned an error", scanner.Err())

	err = json.Unmarshal(marshalledData, &UNMARSHALED_DATA)
	f.ReportIfError("Failed to unmarshal json", err)

	for _, user := range UNMARSHALED_DATA {
		if user.UserName == username {
			target = user
		}
	}

	return target.Forename, target.Lists
}

func SaveFile() {
	byteArr, err := json.MarshalIndent(UNMARSHALED_DATA, "  ", "  ")
	f.ReportIfError("Failed to marshal json", err)

	err = os.WriteFile(FILE_PATH, byteArr, 0644)
	f.ReportIfError("Writing your byte array to the file selected failed", err)
}

func GetLists() []t.List {
	var lists []t.List

	for _, user := range UNMARSHALED_DATA {
		if user.UserName == USERNAME {
			lists = user.Lists
		}
	}

	return lists
}

func GetListItems(listsReadOnly []t.List, listName string) []t.Item {
	var items []t.Item

	for _, list := range listsReadOnly {
		if list.ListName == strings.TrimSpace(listName) {
			items := list.Items
			for _, item := range items {
				complete := f.Turn(item.ItemIsChecked, "Complete", "Incomplete")
				fmt.Printf("    %v : %v : %v\n", item.ItemName, item.ItemDesc, complete)
			}
		}
	}
	return items
}
