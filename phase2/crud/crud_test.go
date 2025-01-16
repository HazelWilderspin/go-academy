package crud

import (
	"phase2/types"
	"testing"

	"github.com/google/uuid"
)

var userId = uuid.Must(uuid.Parse("5ad98d52-5c4e-412e-a6ad-d33598f22c9e"))
var listId = uuid.Must(uuid.Parse("adde3876-594d-4d64-98b0-542b06ed3659"))
var itemId = uuid.Must(uuid.Parse("21ff4827-0126-4eac-a7b7-a343c1dffa22"))

func TestReadUser(t *testing.T) {
	FILE_PATH = "mockData.json"
	expectedForename := "Alice"

	user, err := ReadUser(userId)

	if err != nil {
		t.Errorf("ReadUser produced an error: %v", err)
	}

	if user.Forename != expectedForename {
		t.Errorf("Targets name should be %s not: %s", expectedForename, user.Forename)
	}
}

func TestCreateList(t *testing.T) {
	newList := makeList()

	err := CreateList(userId, newList)
	if err != nil {
		t.Errorf("CreateList produced an error: %v", err)
	}

	err = DeleteList(userId, newList.ListId)
	if err != nil {
		t.Errorf("DeleteList produced an error: %v", err)
	}
}

func TestReadList(t *testing.T) {
	newList := makeList()

	err := CreateList(userId, newList)
	if err != nil {
		t.Errorf("CreateList produced an error: %v", err)
	}

	list, err := ReadList(userId, newList.ListId)
	if err != nil {
		t.Errorf("ReadList produced an error: %v", err)
	}

	if list.ListId != newList.ListId {
		t.Errorf("Target uuid is incorrect: %v", newList.ListId)
	}

	err = DeleteList(userId, newList.ListId)
	if err != nil {
		t.Errorf("DeleteList produced an error: %v", err)
	}
}

func TestUpdateListName(t *testing.T) {
	newList := makeList()

	err := CreateList(userId, newList)
	if err != nil {
		t.Errorf("CreateList produced an error: %v", err)
	}

	expectedNewName := "TEST LIST NAME UPDATE"
	err = UpdateListName(userId, newList.ListId, expectedNewName)
	if err != nil {
		t.Errorf("UpdateListName produced an error: %v", err)
	}

	list, err := ReadList(userId, newList.ListId)
	if err != nil {
		t.Errorf("ReadList produced an error: %v", err)
	}
	if list.ListName != expectedNewName {
		t.Errorf("To pass the List name should have been updated to: %s", list.ListName)
	}

	err = DeleteList(userId, newList.ListId)
	if err != nil {
		t.Errorf("DeleteList produced an error: %v", err)
	}
}

func TestUpdateListToggleCompletion(t *testing.T) {
	newList := makeList()

	err := CreateList(userId, newList)
	if err != nil {
		t.Errorf("CreateList produced an error: %v", err)
	}

	expectedNewValue := true
	err = UpdateListToggleCompletion(userId, newList.ListId, false)
	if err != nil {
		t.Errorf("UpdateListToggleCompletion produced an error: %v", err)
	}

	list, err := ReadList(userId, newList.ListId)
	if err != nil {
		t.Errorf("ReadList produced an error: %v", err)
	}
	if list.IsComplete != expectedNewValue {
		t.Errorf("To pass the List check marker should have been updated to: %t", list.IsComplete)
	}

	err = DeleteList(userId, newList.ListId)
	if err != nil {
		t.Errorf("DeleteList produced an error: %v", err)
	}
}

func TestDeleteList(t *testing.T) {
	newList := makeList()

	err := CreateList(userId, newList)
	if err != nil {
		t.Errorf("CreateList produced an error: %v", err)
	}

	list, err := ReadList(userId, newList.ListId)
	if err != nil {
		t.Errorf("ReadList produced an error: %v", err)
	}
	if list.ListId != newList.ListId {
		t.Errorf("Target uuid is incorrect: %v", newList.ListId)
	}

	err = DeleteList(userId, newList.ListId)
	if err != nil {
		t.Errorf("DeleteList produced an error: %v", err)
	}

	list, err = ReadList(userId, newList.ListId)
	if err == nil {
		t.Errorf("ReadList should proc an error once the target list has been deleted")
	}
}

func makeItem() types.Item {
	newItem := types.Item{
		ItemId:        uuid.New(),
		ItemName:      "ITEM ONE",
		ItemDesc:      "Description",
		ItemIsChecked: false}
	return newItem
}

func makeList() types.List {
	var items []types.Item
	newItem := makeItem()
	items = append(items, newItem)

	newList := types.List{ListId: uuid.New(),
		ListName:   "UNIT TEST CREATE LIST",
		InitDate:   "2012-04-23T18:25:43.511Z",
		IsComplete: false,
		Items:      items}

	return newList
}
