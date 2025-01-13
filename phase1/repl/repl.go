package repl

import (
	"bufio"
	"fmt"
	"os"
	f "phase1/functions"
	t "phase1/types"
	"strings"

	"github.com/google/uuid"
)

func PrintListOptions() {
	fmt.Println("Options: createList | readLists | readList | deleteList | exit ")
}

func PrintItemOptions() {
	fmt.Println("Options: createItem | readItems | updateItem | deleteItem | exit ")
}

func FormNewList() t.List {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter list name to create:")
	newListName, err := reader.ReadString('\n')
	f.ReportIfError("bufio input resulted in an error", err)

	fmt.Println("Define item name to add an item:")
	newItemName, err := reader.ReadString('\n')
	f.ReportIfError("bufio input resulted in an error", err)

	fmt.Println("Add item description:")
	newItemDesc, err := reader.ReadString('\n')
	f.ReportIfError("bufio input resulted in an error", err)

	var items []t.Item
	newItem := t.Item{ItemId: uuid.New(), ItemName: strings.TrimSpace(newItemName), ItemDesc: strings.TrimSpace(newItemDesc), ItemIsChecked: false}
	items = append(items, newItem)
	newList := t.List{ListId: uuid.New(), ListName: strings.TrimSpace(newListName), InitDate: "2025-04-23T18:25:43.511Z", IsComplete: false, Items: items}

	return newList
}

func FormNewItem() t.Item {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter item name:")
	newItemName, err := reader.ReadString('\n')
	f.ReportIfError("bufio input resulted in an error", err)

	fmt.Println("Add item description:")
	newItemDesc, err := reader.ReadString('\n')
	f.ReportIfError("bufio input resulted in an error", err)

	newItem := t.Item{ItemId: uuid.New(), ItemName: strings.TrimSpace(newItemName), ItemDesc: strings.TrimSpace(newItemDesc), ItemIsChecked: false}

	return newItem
}
