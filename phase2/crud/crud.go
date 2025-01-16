package crud

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"os"
	t "phase2/types"

	// "sync"

	"github.com/google/uuid"
)

var FILE_PATH string = "crud/mockData.json"

// add mutex lock
var UNMARSHALED_DATA_STORE []t.User

func loadMockData() error {
	fmt.Println("##### LOADING DATA #####")
	defer fmt.Println("##### LOADING DATA COMPLETE #####")

	var marshalledData []byte

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
	err = json.Unmarshal(marshalledData, &UNMARSHALED_DATA_STORE)
	if err != nil {
		slog.Error("Failed to unmarshal json", "Server Error: ", err)
	}
	return err
}

func ReadUser(userId uuid.UUID) (t.User, error) {
	if len(UNMARSHALED_DATA_STORE) == 0 {
		err := loadMockData()
		if err != nil {
			slog.Error("Mock data failed to load", "Server Error: ", err)
		}
	}
	for _, user := range UNMARSHALED_DATA_STORE {
		if user.UserDetailId == userId {
			return user, nil
		}
	}
	return t.User{}, errors.New("user not found")
}

func CreateList(userId uuid.UUID, newList t.List) error {
	if len(UNMARSHALED_DATA_STORE) == 0 {
		err := loadMockData()
		if err != nil {
			slog.Error("Mock data failed to load", "Server Error: ", err)
		}
	}
	for key, user := range UNMARSHALED_DATA_STORE {
		if user.UserDetailId == userId {
			UNMARSHALED_DATA_STORE[key].Lists = append(UNMARSHALED_DATA_STORE[key].Lists, newList)
			err := saveFile()
			return err
		}
	}
	return errors.New("unable to create list")
}

func ReadList(userId uuid.UUID, listId uuid.UUID) (t.List, error) {
	if len(UNMARSHALED_DATA_STORE) == 0 {
		err := loadMockData()
		if err != nil {
			slog.Error("Mock data failed to load", "Server Error: ", err)
		}
	}
	for _, user := range UNMARSHALED_DATA_STORE {
		if user.UserDetailId == userId {
			for _, list := range user.Lists {
				if list.ListId == listId {
					return list, nil
				}
			}
		}
	}
	return t.List{}, errors.New("list not found")
}

func UpdateListName(userId uuid.UUID, listId uuid.UUID, newName string) error {
	if len(UNMARSHALED_DATA_STORE) == 0 {
		err := loadMockData()
		if err != nil {
			slog.Error("Mock data failed to load", "Server Error: ", err)
		}
	}
	for key1, user := range UNMARSHALED_DATA_STORE {
		if user.UserDetailId == userId {
			for key2, list := range user.Lists {
				if list.ListId == listId {
					UNMARSHALED_DATA_STORE[key1].Lists[key2].ListName = newName
					err := saveFile()
					return err
				}
			}
		}
	}
	return errors.New("update list name unsuccessful")
}

func UpdateListToggleCompletion(userId uuid.UUID, listId uuid.UUID, completion bool) error {
	if len(UNMARSHALED_DATA_STORE) == 0 {
		err := loadMockData()
		if err != nil {
			slog.Error("Mock data failed to load", "Server Error: ", err)
		}
	}
	for key1, user := range UNMARSHALED_DATA_STORE {
		if user.UserDetailId == userId {
			for key2, list := range user.Lists {
				if list.ListId == listId {
					if list.IsComplete {
						UNMARSHALED_DATA_STORE[key1].Lists[key2].IsComplete = false
					} else {
						UNMARSHALED_DATA_STORE[key1].Lists[key2].IsComplete = true
					}
					err := saveFile()
					return err
				}
			}
		}
	}
	return errors.New("update list completion unsuccessful")
}

func DeleteList(userId uuid.UUID, listId uuid.UUID) error {
	if len(UNMARSHALED_DATA_STORE) == 0 {
		err := loadMockData()
		if err != nil {
			slog.Error("Mock data failed to load", "Server Error: ", err)
		}
	}
	for key1, user := range UNMARSHALED_DATA_STORE {
		if user.UserDetailId == userId {
			for key2, list := range user.Lists {
				if list.ListId == listId {
					if key2 >= len(user.Lists) {
						UNMARSHALED_DATA_STORE[key1].Lists = UNMARSHALED_DATA_STORE[key1].Lists[:key2]
						err := saveFile()
						return err
					} else {
						UNMARSHALED_DATA_STORE[key1].Lists = append(UNMARSHALED_DATA_STORE[key1].Lists[:key2], UNMARSHALED_DATA_STORE[key1].Lists[key2+1:]...)
						err := saveFile()
						return err
					}
				}
			}
		}
	}
	return errors.New("list not found")
}

func CreateItem(userId uuid.UUID, listId uuid.UUID, newItem t.Item) error {
	if len(UNMARSHALED_DATA_STORE) == 0 {
		err := loadMockData()
		if err != nil {
			slog.Error("Mock data failed to load", "Server Error: ", err)
		}
	}

	for key1, user := range UNMARSHALED_DATA_STORE {
		if user.UserDetailId == userId {
			for key2, list := range user.Lists {
				if list.ListId == listId {
					UNMARSHALED_DATA_STORE[key1].Lists[key2].Items = append(UNMARSHALED_DATA_STORE[key1].Lists[key2].Items, newItem)
					err := saveFile()
					return err
				}
			}
		}
	}
	return errors.New("unable to create item")
}

func ReadItem(userId uuid.UUID, listId uuid.UUID, itemId uuid.UUID) (t.Item, error) {
	if len(UNMARSHALED_DATA_STORE) == 0 {
		err := loadMockData()
		if err != nil {
			slog.Error("Mock data failed to load", "Server Error: ", err)
		}
	}

	for _, user := range UNMARSHALED_DATA_STORE {
		if user.UserDetailId == userId {
			for _, list := range user.Lists {
				if list.ListId == listId {
					for _, item := range list.Items {
						if item.ItemId == itemId {
							return item, nil
						}
					}
				}
			}
		}
	}
	return t.Item{}, errors.New("item not found")
}

func UpdateItem(userId uuid.UUID, listId uuid.UUID, revisedItem t.Item) (t.Item, error) {
	if len(UNMARSHALED_DATA_STORE) == 0 {
		err := loadMockData()
		if err != nil {
			slog.Error("Mock data failed to load", "Server Error: ", err)
		}
	}

	for key1, user := range UNMARSHALED_DATA_STORE {
		if user.UserDetailId == userId {
			for key2, list := range user.Lists {
				if list.ListId == listId {
					for key3, item := range list.Items {
						if item.ItemId == revisedItem.ItemId {
							UNMARSHALED_DATA_STORE[key1].Lists[key2].Items[key3].ItemName = revisedItem.ItemName
							UNMARSHALED_DATA_STORE[key1].Lists[key2].Items[key3].ItemDesc = revisedItem.ItemDesc
							UNMARSHALED_DATA_STORE[key1].Lists[key2].Items[key3].ItemIsChecked = revisedItem.ItemIsChecked
							err := saveFile()
							return UNMARSHALED_DATA_STORE[key1].Lists[key2].Items[key3], err
						}
					}
				}
			}
		}
	}
	return t.Item{}, errors.New("item not found")
}

func DeleteItem(userId uuid.UUID, listId uuid.UUID, itemId uuid.UUID) error {
	if len(UNMARSHALED_DATA_STORE) == 0 {
		err := loadMockData()
		if err != nil {
			slog.Error("Mock data failed to load", "Server Error: ", err)
		}
	}

	for key1, user := range UNMARSHALED_DATA_STORE {
		if user.UserDetailId == userId {
			for key2, list := range user.Lists {
				if list.ListId == listId {
					for key3, item := range list.Items {
						if item.ItemId == itemId {
							if key3 >= len(list.Items) {
								UNMARSHALED_DATA_STORE[key1].Lists[key2].Items = UNMARSHALED_DATA_STORE[key1].Lists[key2].Items[:key3]
								err := saveFile()
								return err
							} else {
								UNMARSHALED_DATA_STORE[key1].Lists[key2].Items = append(UNMARSHALED_DATA_STORE[key1].Lists[key2].Items[:key3], UNMARSHALED_DATA_STORE[key1].Lists[key2].Items[key3+1:]...)
								err := saveFile()
								return err
							}
						}
					}
				}
			}
		}
	}
	return errors.New("item not found")
}

func saveFile() error {
	byteArr, err := json.MarshalIndent(UNMARSHALED_DATA_STORE, "  ", "  ")
	if err != nil {
		slog.Error("Failed to marshal json", "err", err)
	}
	err = os.WriteFile(FILE_PATH, byteArr, 0644)
	if err != nil {
		slog.Error("Writing your byte array to the file selected failed", "err", err)
	}
	return err
}
