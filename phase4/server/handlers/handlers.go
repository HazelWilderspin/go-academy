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

	var Err error
	defer func() {
		cancelCtx()
		if Err != nil {
			slog.Error(Err.Error(), "trace_id", ctx.Value(traceIdKey))
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(Err.Error()))
		}
	}()

	var unmarshaledBody GetUserRequestBody

	bodyByteArr, err := io.ReadAll(r.Body)
	if err != nil {
		Err = err
		return
	}

	err = json.Unmarshal(bodyByteArr, &unmarshaledBody)
	if err != nil {
		Err = err
		return
	}

	user, err := actor.AddGetUserToRequestChannel(unmarshaledBody.Username, "GetUser")
	if err != nil {
		Err = err
		return
	}

	marshaledUser, err := json.MarshalIndent(&user, "  ", "  ")
	if err != nil {
		Err = err
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

	var Err error
	defer func() {
		cancelCtx()
		if Err != nil {
			slog.Error(Err.Error(), "trace_id", ctx.Value(traceIdKey))
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(Err.Error()))
		}
	}()

	var unmarshaledBody PostListRequestBody

	bodyByteArr, err := io.ReadAll(r.Body)
	if err != nil {
		Err = err
		return
	}

	err = json.Unmarshal(bodyByteArr, &unmarshaledBody)
	if err != nil {
		Err = err
		return
	}

	err = actor.AddPostListToRequestChannel(unmarshaledBody.UserId, unmarshaledBody.NewList, "PostList")
	if err != nil {
		Err = err
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

	var Err error
	defer func() {
		cancelCtx()
		if Err != nil {
			slog.Error(Err.Error(), "trace_id", ctx.Value(traceIdKey))
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(Err.Error()))
		}
	}()

	var unmarshaledBody PutListNameRequestBody

	bodyByteArr, err := io.ReadAll(r.Body)
	if err != nil {
		Err = err
		return
	}

	err = json.Unmarshal(bodyByteArr, &unmarshaledBody)
	if err != nil {
		Err = err
		return
	}

	err = crud.UpdateListName(unmarshaledBody.UserId, unmarshaledBody.ListId, unmarshaledBody.NewListName)
	if err != nil {
		Err = err
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

	var Err error
	defer func() {
		cancelCtx()
		if Err != nil {
			slog.Error(Err.Error(), "trace_id", ctx.Value(traceIdKey))
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(Err.Error()))
		}
	}()

	var unmarshaledBody PutListCompletionRequestBody

	bodyByteArr, err := io.ReadAll(r.Body)
	if err != nil {
		Err = err
		return
	}
	err = json.Unmarshal(bodyByteArr, &unmarshaledBody)
	if err != nil {
		Err = err
		return
	}

	err = crud.UpdateListToggleCompletion(unmarshaledBody.UserId, unmarshaledBody.ListId, unmarshaledBody.ListIsComplete)
	if err != nil {
		Err = err
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

	var Err error
	defer func() {
		cancelCtx()
		if Err != nil {
			slog.Error(Err.Error(), "trace_id", ctx.Value(traceIdKey))
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(Err.Error()))
		}
	}()

	var unmarshaledBody DeleteListRequestBody

	bodyByteArr, err := io.ReadAll(r.Body)
	if err != nil {
		Err = err
		return
	}

	err = json.Unmarshal(bodyByteArr, &unmarshaledBody)
	if err != nil {
		Err = err
		return
	}

	err = actor.AddDeleteListToRequestChannel(unmarshaledBody.UserId, unmarshaledBody.ListId, "DeleteList")
	if err != nil {
		Err = err
		return
	}

	fmt.Println("--------- SUCCESS")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("SUCCESS"))
}

func PostItem(w http.ResponseWriter, r *http.Request) {
	fmt.Println("--- POST ITEM START")
	defer fmt.Println("--- POST ITEM END")

	ctx, cancelCtx := context.WithCancel(r.Context())

	var Err error
	defer func() {
		cancelCtx()
		if Err != nil {
			slog.Error(Err.Error(), "trace_id", ctx.Value(traceIdKey))
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(Err.Error()))
		}
	}()

	var unmarshaledBody PostItemRequestBody

	bodyByteArr, err := io.ReadAll(r.Body)
	if err != nil {
		Err = err
		return
	}

	err = json.Unmarshal(bodyByteArr, &unmarshaledBody)
	if err != nil {
		Err = err
		return
	}

	err = actor.AddPostItemToRequestChannel(unmarshaledBody.UserId, unmarshaledBody.ListId, unmarshaledBody.NewItem, "PostItem")
	if err != nil {
		Err = err
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

	var Err error
	defer func() {
		cancelCtx()
		if Err != nil {
			slog.Error(Err.Error(), "trace_id", ctx.Value(traceIdKey))
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(Err.Error()))
		}
	}()

	var unmarshaledBody PutItemRequestBody

	bodyByteArr, err := io.ReadAll(r.Body)
	if err != nil {
		Err = err
		return
	}

	err = json.Unmarshal(bodyByteArr, &unmarshaledBody)
	if err != nil {
		Err = err
		return
	}

	err = actor.AddPutItemToRequestChannel(unmarshaledBody.UserId, unmarshaledBody.ListId, unmarshaledBody.Item, "PutItem")
	if err != nil {
		Err = err
		return
	}

	fmt.Println("--------- SUCCESS")
	w.WriteHeader(http.StatusOK)
}

func DeleteItem(w http.ResponseWriter, r *http.Request) {
	fmt.Println("--- DELETE ITEM START")
	defer fmt.Println("--- DELETE ITEM END")

	ctx, cancelCtx := context.WithCancel(r.Context())

	var Err error
	defer func() {
		cancelCtx()
		if Err != nil {
			slog.Error(Err.Error(), "trace_id", ctx.Value(traceIdKey))
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(Err.Error()))
		}
	}()

	var unmarshaledBody DeleteItemRequestBody

	bodyByteArr, err := io.ReadAll(r.Body)
	if err != nil {
		Err = err
		return
	}

	err = json.Unmarshal(bodyByteArr, &unmarshaledBody)
	if err != nil {
		Err = err
		return
	}

	err = actor.AddDeleteItemToRequestChannel(unmarshaledBody.UserId, unmarshaledBody.ListId, unmarshaledBody.ItemId, "DeleteItem")
	if err != nil {
		Err = err
		return
	}

	fmt.Println("--------- SUCCESS")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("SUCCESS"))
}
