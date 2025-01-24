package handlers

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"strings"
	"text/template"

	"client/actions"

	"github.com/google/uuid"
)

var DUMMY_CACHE Cache

func HomePageHandler(w http.ResponseWriter, req *http.Request) {
	var err error
	defer func() {
		if err != nil {
			slog.Error(err.Error())
		}
	}()

	page_template, err := template.ParseFiles("html/login.html")
	if err != nil {
		return
	}

	err = page_template.Execute(w, nil)
	if err != nil {
		return
	}
}

func LoginHandler(w http.ResponseWriter, req *http.Request) {
	var err error
	defer func() {
		if err != nil {
			slog.Error(err.Error())
			w.Write([]byte(err.Error()))
		}
	}()

	username := strings.TrimSpace(strings.ToUpper(req.FormValue("user_name")))

	reqBody, err := json.Marshal(GetUserRequestBody{username})
	if err != nil {
		return
	}

	marshalledData, err := actions.GetUser(reqBody)
	if err != nil {
		return
	}

	var user User
	err = json.Unmarshal(marshalledData, &user)
	if err != nil {
		return
	}

	page_template, err := template.ParseFiles("html/myLists.html")
	if err != nil {
		return
	}

	DUMMY_CACHE = Cache{
		UserDetailId: user.UserDetailId,
		Username:     user.UserName,
		Forename:     user.Forename,
		ListCount:    len(user.Lists),
		Lists:        user.Lists,
	}

	err = page_template.Execute(w, DUMMY_CACHE)
	if err != nil {
		return
	}
}

func NewListPageHandler(w http.ResponseWriter, req *http.Request) {
	var err error
	defer func() {
		if err != nil {
			slog.Error(err.Error())
		}
	}()

	page_template, err := template.ParseFiles("html/newList.html")
	if err != nil {
		return
	}
	err = page_template.Execute(w, DUMMY_CACHE)
	if err != nil {
		return
	}
}

func SubmitListFormHandler(w http.ResponseWriter, req *http.Request) {
	var err error
	defer func() {
		if err != nil {
			slog.Error(err.Error())
		}
	}()

	var items []Item
	newItem1 := makeItem(req.FormValue("item_name_1"), req.FormValue("item_desc_1"))
	newItem2 := makeItem(req.FormValue("item_name_2"), req.FormValue("item_desc_2"))
	newItem3 := makeItem(req.FormValue("item_name_3"), req.FormValue("item_desc_3"))
	items = append(items, newItem1, newItem2, newItem3)
	newList := makeList(req.FormValue("list_name"), items)

	reqBody, err := json.Marshal(PostListRequestBody{DUMMY_CACHE.UserDetailId, newList})
	if err != nil {
		return
	}

	err = actions.PostList(reqBody)
	if err != nil {
		return
	}

	err = refreshCache()
	if err != nil {
		return
	}

	page_template, err := template.ParseFiles("html/myLists.html")
	if err != nil {
		return
	}

	err = page_template.Execute(w, DUMMY_CACHE)
	if err != nil {
		return
	}
}

func DeleteListHandler(w http.ResponseWriter, req *http.Request) {
	var err error
	defer func() {
		if err != nil {
			slog.Error(err.Error())
		}
	}()

	listId := uuid.Must(uuid.Parse(req.FormValue("list_delete_btn")))
	reqBody, err := json.Marshal(DeleteListRequestBody{DUMMY_CACHE.UserDetailId, listId})
	if err != nil {
		return
	}

	err = actions.DeleteList(reqBody)
	if err != nil {
		return
	}

	err = refreshCache()
	if err != nil {
		return
	}

	page_template, err := template.ParseFiles("html/myLists.html")
	if err != nil {
		return
	}

	err = page_template.Execute(w, DUMMY_CACHE)
	if err != nil {
		return
	}
}

func AddItemHandler(w http.ResponseWriter, req *http.Request) {
	var err error
	defer func() {
		if err != nil {
			slog.Error(err.Error())
		}
	}()

	newItem := makeItem(req.FormValue("new_item_name"), req.FormValue("new_item_desc"))
	listId := uuid.Must(uuid.Parse(req.FormValue("item_add_btn")))

	reqBody, err := json.Marshal(AddItemRequestBody{DUMMY_CACHE.UserDetailId, listId, newItem})
	if err != nil {
		return
	}

	err = actions.PostItem(reqBody)
	if err != nil {
		return
	}

	err = refreshCache()
	if err != nil {
		return
	}

	page_template, err := template.ParseFiles("html/myLists.html")
	if err != nil {
		return
	}

	err = page_template.Execute(w, DUMMY_CACHE)
	if err != nil {
		return
	}
}

func DeleteItemHandler(w http.ResponseWriter, req *http.Request) {
	var err error
	defer func() {
		if err != nil {
			slog.Error(err.Error())
		}
	}()

	itemId := uuid.Must(uuid.Parse(req.FormValue("item_delete_btn")))
	listId := uuid.Must(uuid.Parse(req.FormValue("list_id")))

	reqBody, err := json.Marshal(DeleteItemRequestBody{DUMMY_CACHE.UserDetailId, listId, itemId})
	if err != nil {
		return
	}

	err = actions.DeleteItem(reqBody)
	if err != nil {
		return
	}

	err = refreshCache()
	if err != nil {
		return
	}

	page_template, err := template.ParseFiles("html/myLists.html")
	if err != nil {
		return
	}

	err = page_template.Execute(w, DUMMY_CACHE)
	if err != nil {
		return
	}
}

// Hacked rubbish so I don't need to learn alpine form binding!
func refreshCache() error {
	reqBody, err := json.Marshal(GetUserRequestBody{DUMMY_CACHE.Username})
	if err != nil {
		return err
	}

	marshalledData, err := actions.GetUser(reqBody)
	if err != nil {
		return err
	}

	var user User
	err = json.Unmarshal(marshalledData, &user)
	if err != nil {
		return err
	}

	DUMMY_CACHE = Cache{
		UserDetailId: user.UserDetailId,
		Username:     user.UserName,
		Forename:     user.Forename,
		ListCount:    len(user.Lists),
		Lists:        user.Lists,
	}

	return err
}
