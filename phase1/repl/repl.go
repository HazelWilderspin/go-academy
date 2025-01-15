package repl

import (
	"bufio"
	"fmt"
	"os"
	types "phase1/types"
	"strings"

	"github.com/google/uuid"
)

func PrintListOptions() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Options: createList | readLists | readAndUpdateList | deleteList | exit ")
	command, err := reader.ReadString('\n')

	// err = fmt.Errorf("Errrrrroooroororor")
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(command), err
}

func PrintItemOptions() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Options: showItems | addItem | toggleItemCheck | reassignItemName | editItemDescription | deleteItem | back ")
	command, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	reader.Reset(reader)
	return strings.TrimSpace(command), err
}

func LogIn() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter your username:")
	username, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	reader.Reset(reader)
	return username, err
}

func Welcome(name string, lists []types.List) {
	fmt.Printf("Hello %s \n You have %d list(s), their names are as follows: \n", name, len(lists))
	for _, value := range lists {
		fmt.Printf("    %v\n", value.ListName)
	}
}

func FormNewList() (types.List, error) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter list name to create:")
	newListName, err := reader.ReadString('\n')
	if err != nil {
		return types.List{}, err
	}

	fmt.Println("Define item name to add an item:")
	newItemName, err := reader.ReadString('\n')
	if err != nil {
		return types.List{}, err
	}

	fmt.Println("Add item description:")
	newItemDesc, err := reader.ReadString('\n')
	if err != nil {
		return types.List{}, err
	}

	var items []types.Item

	newItem := types.Item{
		ItemId:        uuid.New(),
		ItemName:      strings.TrimSpace(newItemName),
		ItemDesc:      strings.TrimSpace(newItemDesc),
		ItemIsChecked: false}

	items = append(items, newItem)

	newList := types.List{
		ListId:     uuid.New(),
		ListName:   strings.TrimSpace(newListName),
		InitDate:   "2025-04-23T18:25:43.511Z",
		IsComplete: false,
		Items:      items}

	reader.Reset(reader)
	return newList, err
}

func FormUpdateList() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter list name to read:")
	listName, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	reader.Reset(reader)
	return strings.TrimSpace(listName), err
}

func FormListToDelete(listsReadOnly []types.List) (string, error) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("You have %d list(s), their names are as follows: \n", len(listsReadOnly))
	for _, value := range listsReadOnly {
		fmt.Printf("    %v\n", value.ListName)
	}

	fmt.Println("Enter list name to delete it:")
	listName, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	reader.Reset(reader)
	return strings.TrimSpace(listName), err
}

func FormNewItem() (types.Item, error) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter item name:")
	newItemName, err := reader.ReadString('\n')
	if err != nil {
		return types.Item{}, err
	}

	fmt.Println("Add item description:")
	newItemDesc, err := reader.ReadString('\n')
	if err != nil {
		return types.Item{}, err
	}

	newItem := types.Item{
		ItemId:        uuid.New(),
		ItemName:      strings.TrimSpace(newItemName),
		ItemDesc:      strings.TrimSpace(newItemDesc),
		ItemIsChecked: false}
	reader.Reset(reader)
	return newItem, err
}

func FormItemNameToToggleCheck() (string, error) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter item name to toggle check status:")
	itemName, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	reader.Reset(reader)
	return strings.TrimSpace(itemName), err
}

func FormItemNameToUpdate() (string, error) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter name of target item:")
	newItemName, err := reader.ReadString('\n')

	if err != nil {
		return "", err
	}
	reader.Reset(reader)
	return strings.TrimSpace(newItemName), err
}

func FormNewItemName() (string, error) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter a new name for your item:")
	newItemName, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	reader.Reset(reader)
	return strings.TrimSpace(newItemName), err
}

func FormNewItemDesc() (string, error) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Give your item a new description:")
	itemDesc, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	reader.Reset(reader)
	return strings.TrimSpace(itemDesc), err
}
