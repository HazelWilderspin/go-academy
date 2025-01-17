package crud

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/google/uuid"
)

var FILE_PATH string = "crud/mockData.json"
var STORE Store

func loadMockData() error {
	fmt.Println("------ LOADING DATA")
	defer fmt.Println("------ LOADING DATA COMPLETE")

	var marshalledData []byte

	file, err := os.Open(FILE_PATH)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		marshalledData = append(marshalledData, scanner.Bytes()...)
	}
	if scanner.Err() != nil {
		return scanner.Err()
	}
	err = json.Unmarshal(marshalledData, &STORE.Data)
	if err != nil {
		return err
	}
	return err
}

func ReadUser(userId uuid.UUID) (User, error) {

	readUserLock := STORE.Lock.RLocker()
	readUserLock.Lock()
	defer readUserLock.Unlock()

	if len(STORE.Data) == 0 {
		err := loadMockData()
		if err != nil {
			return User{}, err
		}
	}

	for _, user := range STORE.Data {
		if user.UserDetailId == userId {
			return user, nil
		}
	}

	return User{}, errors.New("user not found")
}

func CreateList(userId uuid.UUID, newList List) error {

	STORE.Lock.Lock()
	defer STORE.Lock.Unlock()

	if len(STORE.Data) == 0 {
		err := loadMockData()
		if err != nil {
			return err
		}
	}

	for key, user := range STORE.Data {
		if user.UserDetailId == userId {
			STORE.Data[key].Lists = append(STORE.Data[key].Lists, newList)
			err := saveFile()
			return err
		}
	}

	return errors.New("unable to create list")
}

func ReadList(userId uuid.UUID, listId uuid.UUID) (List, error) {

	readListLock := STORE.Lock.RLocker()
	readListLock.Lock()
	defer readListLock.Unlock()

	if len(STORE.Data) == 0 {
		err := loadMockData()
		if err != nil {
			return List{}, err
		}
	}

	for _, user := range STORE.Data {
		if user.UserDetailId == userId {
			for _, list := range user.Lists {
				if list.ListId == listId {
					return list, nil
				}
			}
		}
	}

	return List{}, errors.New("list not found")
}

func UpdateListName(userId uuid.UUID, listId uuid.UUID, newName string) error {

	STORE.Lock.Lock()
	defer STORE.Lock.Unlock()

	if len(STORE.Data) == 0 {
		err := loadMockData()
		if err != nil {
			return err
		}
	}

	for key1, user := range STORE.Data {
		if user.UserDetailId == userId {
			for key2, list := range user.Lists {
				if list.ListId == listId {
					STORE.Data[key1].Lists[key2].ListName = newName
					err := saveFile()
					return err
				}
			}
		}
	}

	return errors.New("update list name unsuccessful")
}

func UpdateListToggleCompletion(userId uuid.UUID, listId uuid.UUID, completion bool) error {

	STORE.Lock.Lock()
	defer STORE.Lock.Unlock()

	if len(STORE.Data) == 0 {
		err := loadMockData()
		if err != nil {
			return err
		}
	}

	for key1, user := range STORE.Data {
		if user.UserDetailId == userId {
			for key2, list := range user.Lists {
				if list.ListId == listId {
					if list.IsComplete {
						STORE.Data[key1].Lists[key2].IsComplete = false
					} else {
						STORE.Data[key1].Lists[key2].IsComplete = true
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

	STORE.Lock.Lock()
	defer STORE.Lock.Unlock()

	if len(STORE.Data) == 0 {
		err := loadMockData()
		if err != nil {
			return err
		}
	}

	for key1, user := range STORE.Data {
		if user.UserDetailId == userId {
			for key2, list := range user.Lists {
				if list.ListId == listId {
					if key2 >= len(user.Lists) {
						STORE.Data[key1].Lists = STORE.Data[key1].Lists[:key2]
						err := saveFile()
						return err
					} else {
						STORE.Data[key1].Lists = append(STORE.Data[key1].Lists[:key2], STORE.Data[key1].Lists[key2+1:]...)
						err := saveFile()
						return err
					}
				}
			}
		}
	}

	return errors.New("list not found")
}

func CreateItem(userId uuid.UUID, listId uuid.UUID, newItem Item) error {

	STORE.Lock.Lock()
	defer STORE.Lock.Unlock()

	if len(STORE.Data) == 0 {
		err := loadMockData()
		if err != nil {
			return err
		}
	}

	for key1, user := range STORE.Data {
		if user.UserDetailId == userId {
			for key2, list := range user.Lists {
				if list.ListId == listId {
					STORE.Data[key1].Lists[key2].Items = append(STORE.Data[key1].Lists[key2].Items, newItem)
					err := saveFile()
					return err
				}
			}
		}
	}

	return errors.New("unable to create item")
}

func ReadItem(userId uuid.UUID, listId uuid.UUID, itemId uuid.UUID) (Item, error) {

	readItemLock := STORE.Lock.RLocker()
	readItemLock.Lock()
	defer readItemLock.Unlock()

	if len(STORE.Data) == 0 {
		err := loadMockData()
		if err != nil {
			return Item{}, err
		}
	}

	for _, user := range STORE.Data {
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

	return Item{}, errors.New("item not found")
}

func UpdateItem(userId uuid.UUID, listId uuid.UUID, revisedItem Item) (Item, error) {
	STORE.Lock.Lock()
	defer STORE.Lock.Unlock()

	if len(STORE.Data) == 0 {
		err := loadMockData()
		if err != nil {
			return Item{}, err
		}
	}

	for key1, user := range STORE.Data {
		if user.UserDetailId == userId {
			for key2, list := range user.Lists {
				if list.ListId == listId {
					for key3, item := range list.Items {
						if item.ItemId == revisedItem.ItemId {
							STORE.Data[key1].Lists[key2].Items[key3].ItemName = revisedItem.ItemName
							STORE.Data[key1].Lists[key2].Items[key3].ItemDesc = revisedItem.ItemDesc
							STORE.Data[key1].Lists[key2].Items[key3].ItemIsChecked = revisedItem.ItemIsChecked
							err := saveFile()
							return STORE.Data[key1].Lists[key2].Items[key3], err
						}
					}
				}
			}
		}
	}

	return Item{}, errors.New("item not found")
}

func DeleteItem(userId uuid.UUID, listId uuid.UUID, itemId uuid.UUID) error {

	STORE.Lock.Lock()
	defer STORE.Lock.Unlock()

	if len(STORE.Data) == 0 {
		err := loadMockData()
		if err != nil {
			return err
		}
	}

	for key1, user := range STORE.Data {
		if user.UserDetailId == userId {
			for key2, list := range user.Lists {
				if list.ListId == listId {
					for key3, item := range list.Items {
						if item.ItemId == itemId {
							if key3 >= len(list.Items) {
								STORE.Data[key1].Lists[key2].Items = STORE.Data[key1].Lists[key2].Items[:key3]
								err := saveFile()
								return err
							} else {
								STORE.Data[key1].Lists[key2].Items = append(STORE.Data[key1].Lists[key2].Items[:key3], STORE.Data[key1].Lists[key2].Items[key3+1:]...)
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

	byteArr, err := json.MarshalIndent(STORE.Data, "  ", "  ")
	if err != nil {
		return err
	}

	err = os.WriteFile(FILE_PATH, byteArr, 0644)
	if err != nil {
		return err
	}

	return err
}
