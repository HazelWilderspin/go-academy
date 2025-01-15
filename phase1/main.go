package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log/slog"
	"os"
	repl "phase1/repl"
	types "phase1/types"
	util "phase1/util"
	"strings"
)

const FILE_PATH string = "mockData.json"

var UNMARSHALED_DATA []types.User
var USERNAME string

func main() {
	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{AddSource: true})))
	slog.Info("Starting")

	username, err := repl.LogIn()
	if err != nil {
		slog.Error("Bufio reader failed", "err", err)
	}
	USERNAME = strings.TrimSpace(username)

	name, lists, err := reloadData()
	if err != nil {
		slog.Error("Reload failed", "err", err)
	}
	repl.Welcome(name, lists)

	listActions()
}

func listActions() {
	command, err := repl.PrintListOptions()
	if err != nil {
		slog.Error("Bufio reader failed test1")
	}

	var listsReadOnly []types.List

	for _, user := range UNMARSHALED_DATA {
		if user.UserName == USERNAME {
			listsReadOnly = user.Lists
		}
	}

	switch strings.TrimSpace(command) {
	case "createList":
		newList, err := repl.FormNewList()
		if err != nil {
			slog.Error("Bufio reader failed", "err", err)
		}

		for key, user := range UNMARSHALED_DATA {
			if user.UserName == USERNAME {
				UNMARSHALED_DATA[key].Lists = append(UNMARSHALED_DATA[key].Lists, newList)
			}
		}
		saveFile()
		reloadData()

	case "readLists":
		fmt.Printf("You have %d list(s), their names are as follows: \n", len(listsReadOnly))

		for _, value := range listsReadOnly {
			fmt.Printf("    %v\n", value.ListName)
		}

	case "readAndUpdateList":
		listName, err := repl.FormUpdateList()
		if err != nil {
			slog.Error("Bufio reader failed", "err", err)
		}

		for _, list := range listsReadOnly {
			if list.ListName == strings.TrimSpace(listName) {
				items := list.Items
				for _, item := range items {
					complete := util.Turn(item.ItemIsChecked, "Complete", "Incomplete")
					fmt.Printf("    %v : %v : %v\n", item.ItemName, item.ItemDesc, complete)
				}
			}
		}
		itemActions(listName)

	case "deleteList":
		listName, err := repl.FormListToDelete(listsReadOnly)
		if err != nil {
			slog.Error("Bufio reader failed", "err", err)
		}

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
		saveFile()
		reloadData()

	case "exit":
		slog.Info("Shutting down")
		fmt.Println("Goodbye")
		return
	default:
		slog.Info("Bad input shutting down")
		return
	}
	listActions()
}

func itemActions(listName string) {
	command, err := repl.PrintItemOptions()
	if err != nil {
		slog.Error("Bufio reader failed", "err", err)
	}

	switch strings.TrimSpace(command) {
	case "showItems":
		for key1, user := range UNMARSHALED_DATA {
			if user.UserName == USERNAME {
				for _, list := range UNMARSHALED_DATA[key1].Lists {
					if list.ListName == strings.TrimSpace(listName) {
						items := list.Items
						for _, item := range items {
							complete := util.Turn(item.ItemIsChecked, "Complete", "Incomplete")
							fmt.Printf("    %v : %v : %v\n", item.ItemName, item.ItemDesc, complete)
						}
					}
				}
			}
		}

	case "addItem":
		newItem, err := repl.FormNewItem()
		if err != nil {
			slog.Error("Bufio reader failed", "err", err)
		}

		for key1, user := range UNMARSHALED_DATA {
			if user.UserName == USERNAME {
				for key2, list := range UNMARSHALED_DATA[key1].Lists {
					if list.ListName == listName {
						UNMARSHALED_DATA[key1].Lists[key2].Items = append(UNMARSHALED_DATA[key1].Lists[key2].Items, newItem)
					}
				}
			}
		}

	case "toggleItemCheck":
		itemName, err := repl.FormItemNameToToggleCheck()
		if err != nil {
			slog.Error("Bufio reader failed", "err", err)
		}

		for key1, user := range UNMARSHALED_DATA {
			if user.UserName == USERNAME {
				for key2, list := range UNMARSHALED_DATA[key1].Lists {
					if list.ListName == listName {
						for key3, item := range UNMARSHALED_DATA[key1].Lists[key2].Items {
							if item.ItemName == itemName {
								inverseCheckVal := util.Turn(item.ItemIsChecked, false, true)
								UNMARSHALED_DATA[key1].Lists[key2].Items[key3].ItemIsChecked = inverseCheckVal
							}
						}
					}
				}
			}
		}

	case "reassignItemName":
		itemName, err := repl.FormItemNameToUpdate()
		if err != nil {
			slog.Error("Bufio reader failed", "err", err)
		}
		newItemName, err := repl.FormNewItemName()
		if err != nil {
			slog.Error("Bufio reader failed", "err", err)
		}

		for key1, user := range UNMARSHALED_DATA {
			if user.UserName == USERNAME {
				for key2, list := range UNMARSHALED_DATA[key1].Lists {
					if list.ListName == listName {
						for key3, item := range UNMARSHALED_DATA[key1].Lists[key2].Items {
							if item.ItemName == itemName {
								UNMARSHALED_DATA[key1].Lists[key2].Items[key3].ItemName = newItemName
							}
						}
					}
				}
			}
		}

	case "editItemDescription":
		itemName, err := repl.FormItemNameToUpdate()
		if err != nil {
			slog.Error("Bufio reader failed", "err", err)
		}

		newDesc, err := repl.FormNewItemDesc()
		if err != nil {
			slog.Error("Bufio reader failed", "err", err)
		}

		for key1, user := range UNMARSHALED_DATA {
			if user.UserName == USERNAME {
				for key2, list := range UNMARSHALED_DATA[key1].Lists {
					if list.ListName == listName {
						for key3, item := range UNMARSHALED_DATA[key1].Lists[key2].Items {
							if item.ItemName == itemName {
								UNMARSHALED_DATA[key1].Lists[key2].Items[key3].ItemDesc = newDesc
							}
						}
					}
				}
			}
		}

	case "deleteItem":
		itemName, err := repl.FormItemNameToUpdate()
		if err != nil {
			slog.Error("Bufio reader failed", "err", err)
		}

		for key1, user := range UNMARSHALED_DATA {
			if user.UserName == USERNAME {
				for key2, list := range UNMARSHALED_DATA[key1].Lists {
					if list.ListName == listName {
						for key3, item := range UNMARSHALED_DATA[key1].Lists[key2].Items {
							if item.ItemName == itemName {
								if key2 >= len(list.Items) {
									UNMARSHALED_DATA[key1].Lists[key2].Items = UNMARSHALED_DATA[key1].Lists[key2].Items[:key3]
								} else {
									UNMARSHALED_DATA[key1].Lists[key2].Items = append(UNMARSHALED_DATA[key1].Lists[key2].Items[:key3], UNMARSHALED_DATA[key1].Lists[key2].Items[key3+1:]...)
								}
							}
						}
					}
				}
			}
		}
	case "back":
		listActions()
		return
	default:
		fmt.Println("User failed out, try something else")
	}

	itemActions(listName)
}

func saveFile() error {
	byteArr, err := json.MarshalIndent(UNMARSHALED_DATA, "  ", "  ")
	if err != nil {
		slog.Error("Failed to marshal json", "err", err)
	}

	err = os.WriteFile(FILE_PATH, byteArr, 0644)
	if err != nil {
		slog.Error("Writing your byte array to the file selected failed", "err", err)
	}

	return err
}

func reloadData() (string, []types.List, error) {
	var marshalledData []byte
	var target types.User

	file, err := os.Open(FILE_PATH)
	if err != nil {
		slog.Error("Opening file was unsuccessful", "err", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		marshalledData = append(marshalledData, scanner.Bytes()...)
	}
	if err != nil {
		slog.Error("bufio scanner returned an error", "err", scanner.Err())
	}

	err = json.Unmarshal(marshalledData, &UNMARSHALED_DATA)
	if err != nil {
		slog.Error("Failed to unmarshal json", "err", err)
	}

	for _, user := range UNMARSHALED_DATA {
		if user.UserName == USERNAME {
			target = user
		}
	}

	return target.Forename, target.Lists, err
}
