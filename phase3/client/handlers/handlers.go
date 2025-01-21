package handlers

import (
	"log/slog"
	"net/http"
	"strings"
	"text/template"

	r "client/actions"

	"github.com/google/uuid"
)

type Store struct {
	UserDetailId uuid.UUID
	Username     string
	Forename     string
	ListCount    int
	Lists        []r.List
}

var STORE Store

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
		}
	}()

	username := req.FormValue("user_name")

	user, err := r.GetUser(strings.TrimSpace(strings.ToUpper(username)))
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	page_template, err := template.ParseFiles("html/myLists.html")
	if err != nil {
		return
	}

	STORE = Store{
		UserDetailId: user.UserDetailId,
		Username:     user.UserName,
		Forename:     user.Forename,
		ListCount:    len(user.Lists),
		Lists:        user.Lists,
	}

	err = page_template.Execute(w, STORE)
	if err != nil {
		return
	}
}

func NewListHandler(w http.ResponseWriter, req *http.Request) {
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

	err = page_template.Execute(w, STORE)
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

	newList := makeList()

	responseCode, err := r.PostList(&STORE.UserDetailId, newList)
	if err != nil {
		slog.Error(err.Error())
		return
	}

	http.Redirect(w, req, "/myLists", responseCode)

}

func makeItem() r.Item {
	newItem := r.Item{
		ItemId:        uuid.New(),
		ItemName:      "ITEM",
		ItemDesc:      "Description",
		ItemIsChecked: false}
	return newItem
}

func makeList() r.List {
	var items []r.Item
	newItem := makeItem()
	items = append(items, newItem)

	newList := r.List{
		ListId:     uuid.New(),
		ListName:   "PHASE 3 TEST CREATE LIST",
		InitDate:   "2012-04-23T18:25:43.511Z",
		IsComplete: false,
		Items:      items}

	return newList
}
