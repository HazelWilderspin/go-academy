package handlers

import (
	"strings"

	"github.com/google/uuid"
)

func makeItem(name string, desc string) Item {
	newItem := Item{
		ItemId:        uuid.New(),
		ItemName:      strings.TrimSpace(name),
		ItemDesc:      strings.TrimSpace(desc),
		ItemIsChecked: false}
	return newItem
}

func makeList(name string, items []Item) List {
	newList := List{
		ListId:     uuid.New(),
		ListName:   strings.TrimSpace(name),
		InitDate:   "2012-04-23T18:25:43.511Z",
		IsComplete: false,
		Items:      items}
	return newList
}
