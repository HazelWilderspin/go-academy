package crud

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/uuid"
)

var userId = uuid.Must(uuid.Parse("5ad98d52-5c4e-412e-a6ad-d33598f22c9e"))
var listId = uuid.Must(uuid.Parse("17f0603d-a086-41f3-b42c-d4021f900431"))
var itemId = uuid.Must(uuid.Parse("4bd26263-efbb-49d6-92f7-48fc0de092fe"))

func makeItem() Item {
	newItem := Item{
		ItemId:        uuid.New(),
		ItemName:      "ITEM ONE",
		ItemDesc:      "Description",
		ItemIsChecked: false}
	return newItem
}

func makeList() List {
	var items []Item
	newItem := makeItem()
	items = append(items, newItem)

	newList := List{ListId: uuid.New(),
		ListName:   "UNIT TEST CREATE LIST",
		InitDate:   "2012-04-23T18:25:43.511Z",
		IsComplete: false,
		Items:      items}

	return newList
}

func trace(name string) func() {
	fmt.Printf("%s --------- entered", name)
	return func() {
		fmt.Printf("%s --------- returned", name)
	}
}

func TestReadUser(t *testing.T) {
	defer trace("TestReadUser")()
	t.Parallel()

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
	defer trace("TestCreateList")()
	t.Parallel()

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
	defer trace("TestReadList")()
	t.Parallel()

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
	defer trace("TestUpdateListName")()
	t.Parallel()

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
	defer trace("TestUpdateListToggleCompletion")()
	t.Parallel()

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
	defer trace("TestDeleteList")()
	t.Parallel()

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

func TestCreateItem(t *testing.T) {
	defer trace("TestCreateItem")()
	t.Parallel()

	FILE_PATH = "mockData.json"

	newItem := makeItem()

	type args struct {
		userId  uuid.UUID
		listId  uuid.UUID
		newItem Item
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "Create a new item on an existing list",
			args:    args{userId, listId, newItem},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CreateItem(tt.args.userId, tt.args.listId, tt.args.newItem); (err != nil) != tt.wantErr {
				t.Errorf("CreateItem() error = %v, wantErr %v", err, tt.wantErr)
			}

		})

	}
	if err := DeleteItem(userId, listId, newItem.ItemId); err != nil {
		t.Errorf("DeleteItem() error = %v", err)
	}
}

func TestReadItem(t *testing.T) {
	defer trace("TestReadItem")()
	t.Parallel()

	FILE_PATH = "mockData.json"

	type args struct {
		userId uuid.UUID
		listId uuid.UUID
		itemId uuid.UUID
	}
	tests := []struct {
		name    string
		args    args
		want    Item
		wantErr bool
	}{
		{
			name:    "Read an existing item on an existing list",
			args:    args{userId, listId, itemId},
			want:    Item{itemId, "Check doors", "Sneak out for a smoke", false},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ReadItem(tt.args.userId, tt.args.listId, tt.args.itemId)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadItem() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadItem() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpdateItem(t *testing.T) {
	defer trace("TestUpdateItem")()
	t.Parallel()

	FILE_PATH = "mockData.json"

	item := makeItem()
	CreateItem(userId, listId, item)
	item.ItemName = "Revised items name"
	item.ItemDesc = "New description"
	item.ItemIsChecked = true

	type args struct {
		userId      uuid.UUID
		listId      uuid.UUID
		revisedItem Item
	}
	tests := []struct {
		name    string
		args    args
		want    Item
		wantErr bool
	}{
		{
			name:    "Update the values on a new item on an existing list",
			args:    args{userId, listId, item},
			want:    Item{item.ItemId, "Revised items name", "New description", true},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			oldItem, err := ReadItem(tt.args.userId, tt.args.listId, item.ItemId)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadItem() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			got, err := UpdateItem(tt.args.userId, tt.args.listId, tt.args.revisedItem)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateItem() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Updated item DeepEqual = %v, want %v", got, tt.want)
				return
			}
			if !reflect.DeepEqual(oldItem.ItemId, tt.want.ItemId) {
				t.Errorf("Old item ID DeepEqual = %v, want %v", oldItem.ItemId, tt.want.ItemId)
				return
			}
			if !cmp.Equal(got, tt.want) {
				t.Errorf("Updated item cmp.Equal = %v, want %v", got, tt.want)
				return
			}
			if !cmp.Equal(oldItem.ItemId, tt.want.ItemId) {
				t.Errorf("Old item ID cmp.Equal = %v, want %v", oldItem.ItemId, tt.want.ItemId)
				return
			}
		})
	}
	if err := DeleteItem(userId, listId, item.ItemId); err != nil {
		t.Errorf("DeleteItem() error = %v", err)
		return
	}
}

func TestDeleteItem(t *testing.T) {
	defer trace("TestDeleteItem")()
	t.Parallel()

	FILE_PATH = "mockData.json"

	item := makeItem()
	CreateItem(userId, listId, item)

	type args struct {
		userId uuid.UUID
		listId uuid.UUID
		itemId uuid.UUID
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "Delete a new item on an existing list",
			args:    args{userId, listId, item.ItemId},
			wantErr: false,
		},
		{
			name:    "Try to delete an item that isn't there",
			args:    args{userId, listId, uuid.Must(uuid.Parse("1ba98ad9-c77f-4f9c-94c3-a2b7e30334ae"))},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := DeleteItem(tt.args.userId, tt.args.listId, tt.args.itemId); (err != nil) != tt.wantErr {
				t.Errorf("DeleteItem() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
