package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"

	"server/actor"
	"server/crud"
)

const (
	traceIdKey CtxKey = "trace_id"
)

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
		}
	}()

	var unmarshaledBody GetUserRequestBody

	bodyByteArr, err := io.ReadAll(r.Body)
	if err != nil {
		return
	}

	err = json.Unmarshal(bodyByteArr, &unmarshaledBody)
	if err != nil {
		return
	}

	user, err := actor.AddGetUserToRequestChannel(unmarshaledBody.Username, "GetUser")
	if err != nil {
		return
	}

	marshaledUser, err := json.MarshalIndent(&user, "  ", "  ")
	if err != nil {
		return
	}

	fmt.Println("--------- SUCCESS")
	w.WriteHeader(http.StatusOK)
	w.Write(marshaledUser)
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
		}
	}()

	var unmarshaledBody PostListRequestBody

	bodyByteArr, err := io.ReadAll(r.Body)
	if err != nil {
		return
	}

	err = json.Unmarshal(bodyByteArr, &unmarshaledBody)
	if err != nil {
		return
	}

	err = actor.AddPostListToRequestChannel(unmarshaledBody.UserId, unmarshaledBody.NewList, "PostList")
	if err != nil {
		return
	}

	fmt.Println("--------- SUCCESS")
	w.WriteHeader(http.StatusOK)
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
		}
	}()

	var unmarshaledBody PutListNameRequestBody

	bodyByteArr, err := io.ReadAll(r.Body)
	if err != nil {
		return
	}

	err = json.Unmarshal(bodyByteArr, &unmarshaledBody)
	if err != nil {
		return
	}

	err = crud.UpdateListName(unmarshaledBody.UserId, unmarshaledBody.ListId, unmarshaledBody.NewListName)
	if err != nil {
		return
	}

	fmt.Println("--------- SUCCESS")
	w.WriteHeader(http.StatusOK)
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
		}
	}()

	var unmarshaledBody PutListCompletionRequestBody

	bodyByteArr, err := io.ReadAll(r.Body)
	if err != nil {
		return
	}

	err = json.Unmarshal(bodyByteArr, &unmarshaledBody)
	if err != nil {
		return
	}

	err = crud.UpdateListToggleCompletion(unmarshaledBody.UserId, unmarshaledBody.ListId, unmarshaledBody.ListIsComplete)
	if err != nil {
		return
	}

	fmt.Println("--------- SUCCESS")
	w.WriteHeader(http.StatusOK)
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
		}
	}()

	var unmarshaledBody DeleteListRequestBody

	bodyByteArr, err := io.ReadAll(r.Body)
	if err != nil {
		return
	}

	err = json.Unmarshal(bodyByteArr, &unmarshaledBody)
	if err != nil {
		return
	}

	err = actor.AddDeleteListToRequestChannel(unmarshaledBody.UserId, unmarshaledBody.ListId, "DeleteList")
	if err != nil {
		return
	}

	fmt.Println("--------- SUCCESS")
	w.WriteHeader(http.StatusOK)
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
		}
	}()

	var unmarshaledBody PostItemRequestBody

	bodyByteArr, err := io.ReadAll(r.Body)
	if err != nil {
		return
	}

	err = json.Unmarshal(bodyByteArr, &unmarshaledBody)
	if err != nil {
		return
	}

	err = actor.AddPostItemToRequestChannel(unmarshaledBody.UserId, unmarshaledBody.ListId, unmarshaledBody.NewItem, "PostItem")
	if err != nil {
		return
	}

	fmt.Println("--------- SUCCESS")
	w.WriteHeader(http.StatusOK)
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
		}
	}()

	var unmarshaledBody PutItemRequestBody

	bodyByteArr, err := io.ReadAll(r.Body)
	if err != nil {
		return
	}

	err = json.Unmarshal(bodyByteArr, &unmarshaledBody)
	if err != nil {
		return
	}

	err = actor.AddPutItemToRequestChannel(unmarshaledBody.UserId, unmarshaledBody.ListId, unmarshaledBody.Item, "PutItem")
	if err != nil {
		return
	}

	fmt.Println("--------- SUCCESS")
	w.WriteHeader(http.StatusOK)
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
		}
	}()

	var unmarshaledBody DeleteItemRequestBody

	bodyByteArr, err := io.ReadAll(r.Body)
	if err != nil {
		return
	}

	err = json.Unmarshal(bodyByteArr, &unmarshaledBody)
	if err != nil {
		return
	}

	err = actor.AddDeleteItemToRequestChannel(unmarshaledBody.UserId, unmarshaledBody.ListId, unmarshaledBody.ItemId, "DeleteItem")
	if err != nil {
		return
	}

	fmt.Println("--------- SUCCESS")
	w.WriteHeader(http.StatusOK)
}
