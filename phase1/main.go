package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/google/uuid"
)

type Item struct {
	itemId             uuid.UUID
	itemName, itemDesc string
	isChecked          bool
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Please enter your username:")
	name, err := reader.ReadString('\n')
	reportIfError(err)

	data, err := getUserData(name)
	fmt.Println("check", data)
	reportIfError(err)
}

func getUserData(userName string) (string, error) {

	options := os.O_RDONLY | os.O_APPEND | os.O_CREATE

	file, err := os.OpenFile("data.json", options, os.FileMode(0600))
	reportIfError(err)
	defer file.Close()

	fmt.Println("check", file)

	decoder := json.NewDecoder(file)

	decoder.Token()

	item := createItem()

	fmt.Println(item)

	// filteredData := []map[string]interface{}{}
	// data := map[string]interface{}{}

	// for decoder.More() {
	// 	decoder.Decode(&data)

	// 	filteredData = append(filteredData, data)
	// }

	return userName, nil
}

func reportIfError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func createItem() Item {
	item := Item{uuid.New(), "New item", "New item description", false}

	return item
}

// func readItem(item *Item) {
// 	item.itemId = int(uuid.New().ID())
// 	item.itemName = "New item"
// 	item.itemDesc = "New item description"
// 	item.isChecked = false
// }

// func updateItem(item *Item) {
// 	item.itemId = int(uuid.New().ID())
// 	item.itemName = "New item"
// 	item.itemDesc = "New item description"
// 	item.isChecked = false
// }

// func deleteItem(item *Item) {
// 	item.itemId = int(uuid.New().ID())
// 	item.itemName = "New item"
// 	item.itemDesc = "New item description"
// 	item.isChecked = false
// }
