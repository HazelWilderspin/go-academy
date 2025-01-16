package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"

	c "phase2/crud"
	t "phase2/types"
)

var USERNAME string

func GetUser(w http.ResponseWriter, req *http.Request) {
	fmt.Println("##### GET USER START #####")
	defer fmt.Println("##### GET USER END #####")

	// ctx, cancel := context.WithCancel(req.Context())
	// ctx = context.WithValue(ctx, traceCtxKey, "123")

	fmt.Println("traceCtxKey", traceCtxKey)
	fmt.Println("record", record)
	fmt.Println("ctx Value traceCtxKey", ctx.Value(traceCtxKey))

	var err error
	defer func() {
		if err != nil {
			cancel()
			slog.Error(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}
	}()

	var unmarshaledBody t.GetUserRequestBody

	bodyByteArr, err := io.ReadAll(req.Body)
	if err != nil {
		err = fmt.Errorf("failed to read request body: %v", err)
		return
	}

	err = json.Unmarshal(bodyByteArr, &unmarshaledBody)
	if err != nil {
		err = fmt.Errorf("failed to unmarshal json: %v", err)
		return
	}

	user, err := c.ReadUser(unmarshaledBody.UserId)
	if err != nil {
		return
	}

	marshaledUser, err := json.MarshalIndent(&user, "  ", "  ")
	if err != nil {
		err = fmt.Errorf("failed to marshal json: %v", err)
	}

	cancel()
	w.WriteHeader(http.StatusOK)
	w.Write(marshaledUser)
}

func GetList(w http.ResponseWriter, req *http.Request) {
	fmt.Println("##### GET LIST START #####")
	defer fmt.Println("##### GET LIST END #####")

	var err error
	defer func() {
		if err != nil {
			slog.Error(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}
	}()

	var unmarshaledBody t.GetListRequestBody

	bodyByteArr, err := io.ReadAll(req.Body)
	if err != nil {
		err = fmt.Errorf("failed to read request body: %v", err)
		return
	}

	err = json.Unmarshal(bodyByteArr, &unmarshaledBody)
	if err != nil {
		err = fmt.Errorf("failed to unmarshal json: %v", err)
		return
	}

	list, err := c.ReadList(unmarshaledBody.UserId, unmarshaledBody.ListId)
	if err != nil {
		return
	}

	marshalledList, err := json.MarshalIndent(&list, "  ", "  ")
	if err != nil {
		err = fmt.Errorf("failed to marshal json: %v", err)
	}
	w.WriteHeader(http.StatusOK)
	w.Write(marshalledList)
}

func PostList(w http.ResponseWriter, req *http.Request) {
	fmt.Println("##### POST LIST START #####")
	defer fmt.Println("##### POST LIST END #####")

	var err error
	defer func() {
		if err != nil {
			slog.Error(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}
	}()

	var unmarshaledBody t.PostListRequestBody

	bodyByteArr, err := io.ReadAll(req.Body)
	if err != nil {
		err = fmt.Errorf("failed to read request body: %v", err)
		return
	}

	err = json.Unmarshal(bodyByteArr, &unmarshaledBody)
	if err != nil {
		err = fmt.Errorf("failed to unmarshal json: %v", err)
		return
	}

	err = c.CreateList(unmarshaledBody.UserId, unmarshaledBody.NewList)
	if err != nil {
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("SUCCESS"))
}

func PutListName(w http.ResponseWriter, req *http.Request) {
	fmt.Println("##### PUT LIST NAME START #####")
	defer fmt.Println("##### PUT LIST NAME END #####")

	var err error
	defer func() {
		if err != nil {
			slog.Error(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}
	}()

	var unmarshaledBody t.PutListNameRequestBody

	bodyByteArr, err := io.ReadAll(req.Body)
	if err != nil {
		err = fmt.Errorf("failed to read request body: %v", err)
		return
	}

	err = json.Unmarshal(bodyByteArr, &unmarshaledBody)
	if err != nil {
		err = fmt.Errorf("failed to unmarshal json: %v", err)
		return
	}

	err = c.UpdateListName(unmarshaledBody.UserId, unmarshaledBody.ListId, unmarshaledBody.NewListName)
	if err != nil {
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("SUCCESS"))
}

func PutListToggleCompletion(w http.ResponseWriter, req *http.Request) {
	fmt.Println("##### TOGGLE LIST COMPLETE START #####")
	defer fmt.Println("##### TOGGLE LIST COMPLETE END #####")

	var err error
	defer func() {
		if err != nil {
			slog.Error(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}
	}()

	var unmarshaledBody t.PutListCompletionRequestBody

	bodyByteArr, err := io.ReadAll(req.Body)
	if err != nil {
		err = fmt.Errorf("failed to read request body: %v", err)
		return
	}
	err = json.Unmarshal(bodyByteArr, &unmarshaledBody)
	if err != nil {
		err = fmt.Errorf("failed to unmarshal json: %v", err)
		return
	}

	err = c.UpdateListToggleCompletion(unmarshaledBody.UserId, unmarshaledBody.ListId, unmarshaledBody.ListIsComplete)
	if err != nil {
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("SUCCESS"))
}

func DeleteList(w http.ResponseWriter, req *http.Request) {
	fmt.Println("##### DELETE LIST START #####")
	defer fmt.Println("##### DELETE LIST END #####")

	var err error
	defer func() {
		if err != nil {
			slog.Error(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}
	}()

	var unmarshaledBody t.DeleteListRequestBody

	bodyByteArr, err := io.ReadAll(req.Body)
	if err != nil {
		err = fmt.Errorf("failed to read request body: %v", err)
		return
	}

	err = json.Unmarshal(bodyByteArr, &unmarshaledBody)
	if err != nil {
		err = fmt.Errorf("failed to unmarshal json: %v", err)
		return
	}

	err = c.DeleteList(unmarshaledBody.UserId, unmarshaledBody.ListId)
	if err != nil {
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("SUCCESS"))
}

func GetItem(w http.ResponseWriter, req *http.Request) {
	fmt.Println("##### GET ITEM START #####")
	defer fmt.Println("##### GET ITEM END #####")

	var err error
	defer func() {
		if err != nil {
			slog.Error(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}
	}()

	var unmarshaledBody t.GetItemRequestBody

	bodyByteArr, err := io.ReadAll(req.Body)
	if err != nil {
		err = fmt.Errorf("failed to read request body: %v", err)
		return
	}

	err = json.Unmarshal(bodyByteArr, &unmarshaledBody)
	if err != nil {
		err = fmt.Errorf("failed to unmarshal json: %v", err)
		return
	}

	item, err := c.ReadItem(unmarshaledBody.UserId, unmarshaledBody.ListId, unmarshaledBody.ItemId)
	if err != nil {
		return
	}

	marshalledItem, err := json.MarshalIndent(&item, "  ", "  ")
	if err != nil {
		err = fmt.Errorf("failed to marshal json: %v", err)
	}
	w.WriteHeader(http.StatusOK)
	w.Write(marshalledItem)

}

func PostItem(w http.ResponseWriter, req *http.Request) {
	fmt.Println("##### POST ITEM START #####")
	defer fmt.Println("##### POST ITEM END #####")

	var err error
	defer func() {
		if err != nil {
			slog.Error(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}
	}()

	var unmarshaledBody t.PostItemRequestBody

	bodyByteArr, err := io.ReadAll(req.Body)
	if err != nil {
		err = fmt.Errorf("failed to read request body: %v", err)
		return
	}

	err = json.Unmarshal(bodyByteArr, &unmarshaledBody)
	if err != nil {
		err = fmt.Errorf("failed to unmarshal json: %v", err)
		return
	}

	err = c.CreateItem(unmarshaledBody.UserId, unmarshaledBody.ListId, unmarshaledBody.NewItem)
	if err != nil {
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("SUCCESS"))
}

func PutItem(w http.ResponseWriter, req *http.Request) {
	fmt.Println("##### PUT ITEM START #####")
	defer fmt.Println("##### PUT ITEM END #####")

	var err error
	defer func() {
		if err != nil {
			slog.Error(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}
	}()

	var unmarshaledBody t.PutItemRequestBody

	bodyByteArr, err := io.ReadAll(req.Body)
	if err != nil {
		err = fmt.Errorf("failed to read request body: %v", err)
		return
	}

	err = json.Unmarshal(bodyByteArr, &unmarshaledBody)
	if err != nil {
		err = fmt.Errorf("failed to unmarshal json: %v", err)
		return
	}

	updatedItem, err := c.UpdateItem(unmarshaledBody.UserId, unmarshaledBody.ListId, unmarshaledBody.Item)
	if err != nil {
		return
	}

	marshalledItem, err := json.MarshalIndent(&updatedItem, "  ", "  ")
	if err != nil {
		err = fmt.Errorf("failed to marshal json: %v", err)
	}

	w.WriteHeader(http.StatusOK)
	w.Write(marshalledItem)
}

func DeleteItem(w http.ResponseWriter, req *http.Request) {
	fmt.Println("##### DELETE ITEM START #####")
	defer fmt.Println("##### DELETE ITEM END #####")

	var err error
	defer func() {
		if err != nil {
			slog.Error(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}
	}()

	var unmarshaledBody t.DeleteItemRequestBody

	bodyByteArr, err := io.ReadAll(req.Body)
	if err != nil {
		err = fmt.Errorf("failed to read request body: %v", err)
		return
	}

	err = json.Unmarshal(bodyByteArr, &unmarshaledBody)
	if err != nil {
		err = fmt.Errorf("failed to unmarshal json: %v", err)
		return
	}

	err = c.DeleteItem(unmarshaledBody.UserId, unmarshaledBody.ListId, unmarshaledBody.ItemId)
	if err != nil {
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("SUCCESS"))
}
