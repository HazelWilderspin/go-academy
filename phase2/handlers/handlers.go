package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"

	c "phase2/crud"
)

const (
	traceIdKey CtxKey = "trace_id"
)

// create a channel to receive a notification on
// use signal.notify to bind an os interrupt to that channel
// listen for interrupt then call context.done

func GetUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("--- GET USER START")
	defer fmt.Println("--- GET USER END")

	ctx, cancelCtx := context.WithCancel(r.Context())

	var err error
	defer func() {
		cancelCtx()
		if err != nil {
			slog.Error(err.Error(), "trace_id", ctx.Value(traceIdKey))
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}
	}()

	var unmarshaledBody GetUserRequestBody

	bodyByteArr, err := io.ReadAll(r.Body)
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

	fmt.Println("--------- SUCCESS")
	w.WriteHeader(http.StatusOK)
	w.Write(marshaledUser)
}

func GetList(w http.ResponseWriter, r *http.Request) {
	fmt.Println("--- GET LIST START")
	defer fmt.Println("--- GET LIST END")

	ctx, cancelCtx := context.WithCancel(r.Context())

	var err error
	defer func() {
		cancelCtx()
		if err != nil {
			slog.Error(err.Error(), "trace_id", ctx.Value(traceIdKey))
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}
	}()

	var unmarshaledBody GetListRequestBody

	bodyByteArr, err := io.ReadAll(r.Body)
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

	fmt.Println("--------- SUCCESS")
	w.WriteHeader(http.StatusOK)
	w.Write(marshalledList)
}

func PostList(w http.ResponseWriter, r *http.Request) {
	fmt.Println("--- POST LIST START")
	defer fmt.Println("--- POST LIST END")

	ctx, cancelCtx := context.WithCancel(r.Context())

	var err error
	defer func() {
		cancelCtx()
		if err != nil {
			slog.Error(err.Error(), "trace_id", ctx.Value(traceIdKey))
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}
	}()

	var unmarshaledBody PostListRequestBody

	bodyByteArr, err := io.ReadAll(r.Body)
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

	fmt.Println("--------- SUCCESS")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("SUCCESS"))
}

func PutListName(w http.ResponseWriter, r *http.Request) {
	fmt.Println("--- PUT LIST NAME START")
	defer fmt.Println("--- PUT LIST NAME END")

	ctx, cancelCtx := context.WithCancel(r.Context())

	var err error
	defer func() {
		cancelCtx()
		if err != nil {
			slog.Error(err.Error(), "trace_id", ctx.Value(traceIdKey))
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}
	}()

	var unmarshaledBody PutListNameRequestBody

	bodyByteArr, err := io.ReadAll(r.Body)
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

	fmt.Println("--------- SUCCESS")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("SUCCESS"))
}

func PutListToggleCompletion(w http.ResponseWriter, r *http.Request) {
	fmt.Println("--- TOGGLE LIST COMPLETE START")
	defer fmt.Println("--- TOGGLE LIST COMPLETE END")

	ctx, cancelCtx := context.WithCancel(r.Context())

	var err error
	defer func() {
		cancelCtx()
		if err != nil {
			slog.Error(err.Error(), "trace_id", ctx.Value(traceIdKey))
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}
	}()

	var unmarshaledBody PutListCompletionRequestBody

	bodyByteArr, err := io.ReadAll(r.Body)
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

	fmt.Println("--------- SUCCESS")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("SUCCESS"))
}

func DeleteList(w http.ResponseWriter, r *http.Request) {
	fmt.Println("--- DELETE LIST START")
	defer fmt.Println("--- DELETE LIST END")

	ctx, cancelCtx := context.WithCancel(r.Context())

	var err error
	defer func() {
		cancelCtx()
		if err != nil {
			slog.Error(err.Error(), "trace_id", ctx.Value(traceIdKey))
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}
	}()

	var unmarshaledBody DeleteListRequestBody

	bodyByteArr, err := io.ReadAll(r.Body)
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

	fmt.Println("--------- SUCCESS")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("SUCCESS"))
}

func GetItem(w http.ResponseWriter, r *http.Request) {
	fmt.Println("--- GET ITEM START")
	defer fmt.Println("--- GET ITEM END")

	ctx, cancelCtx := context.WithCancel(r.Context())

	var err error
	defer func() {
		cancelCtx()
		if err != nil {
			slog.Error(err.Error(), "trace_id", ctx.Value(traceIdKey))
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}
	}()

	var unmarshaledBody GetItemRequestBody

	bodyByteArr, err := io.ReadAll(r.Body)
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

	fmt.Println("--------- SUCCESS")
	w.WriteHeader(http.StatusOK)
	w.Write(marshalledItem)
}

func PostItem(w http.ResponseWriter, r *http.Request) {
	fmt.Println("--- POST ITEM START")
	defer fmt.Println("--- POST ITEM END")

	ctx, cancelCtx := context.WithCancel(r.Context())

	var err error
	defer func() {
		cancelCtx()
		if err != nil {
			slog.Error(err.Error(), "trace_id", ctx.Value(traceIdKey))
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}
	}()

	var unmarshaledBody PostItemRequestBody

	bodyByteArr, err := io.ReadAll(r.Body)
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

	fmt.Println("--------- SUCCESS")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("SUCCESS"))
}

func PutItem(w http.ResponseWriter, r *http.Request) {
	fmt.Println("--- PUT ITEM START")
	defer fmt.Println("--- PUT ITEM END")

	ctx, cancelCtx := context.WithCancel(r.Context())

	var err error
	defer func() {
		cancelCtx()
		if err != nil {
			slog.Error(err.Error(), "trace_id", ctx.Value(traceIdKey))
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}
	}()

	var unmarshaledBody PutItemRequestBody

	bodyByteArr, err := io.ReadAll(r.Body)
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

	fmt.Println("--------- SUCCESS")
	w.WriteHeader(http.StatusOK)
	w.Write(marshalledItem)
}

func DeleteItem(w http.ResponseWriter, r *http.Request) {
	fmt.Println("--- DELETE ITEM START")
	defer fmt.Println("--- DELETE ITEM END")

	ctx, cancelCtx := context.WithCancel(r.Context())

	var err error
	defer func() {
		cancelCtx()
		if err != nil {
			slog.Error(err.Error(), "trace_id", ctx.Value(traceIdKey))
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}
	}()

	var unmarshaledBody DeleteItemRequestBody

	bodyByteArr, err := io.ReadAll(r.Body)
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

	fmt.Println("--------- SUCCESS")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("SUCCESS"))
}
