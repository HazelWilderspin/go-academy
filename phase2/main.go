package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"

	h "phase2/handlers"

	"github.com/google/uuid"
)

type ContextHandler struct {
	slog.Handler
}

func main() {
	slog.SetDefault(slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{AddSource: true})))
	fmt.Printf("-------------------- Go Http server starting --------------------\n")

	mux := http.NewServeMux()

	mux.Handle("/GetUser", middleware(http.HandlerFunc(h.GetUser)))
	mux.Handle("/GetList", middleware(http.HandlerFunc(h.GetList)))
	mux.Handle("/PostList", middleware(http.HandlerFunc(h.PostList)))
	mux.Handle("/PutListName", middleware(http.HandlerFunc(h.PutListName)))
	mux.Handle("/PutListToggleCompletion", middleware(http.HandlerFunc(h.PutListToggleCompletion)))
	mux.Handle("/DeleteList", middleware(http.HandlerFunc(h.DeleteList)))
	mux.Handle("/GetItem", middleware(http.HandlerFunc(h.GetItem)))
	mux.Handle("/PostItem", middleware(http.HandlerFunc(h.PostItem)))
	mux.Handle("/PutItem", middleware(http.HandlerFunc(h.PutItem)))
	mux.Handle("/DeleteItem", middleware(http.HandlerFunc(h.DeleteItem)))

	srv := &http.Server{Addr: "0.0.0.0:8080", Handler: mux}

	err := srv.ListenAndServe()
	if err != nil {
		slog.Error("Bufio reader failed", "Server Error: ", err)
	}
}

func middleware(nextHandler http.Handler) http.Handler {

	midHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var pcs [1]uintptr

		record := slog.NewRecord(time.Now(), slog.Level.Level(1), "test", pcs[0])
		//	gen a trace id to use in ctx?

		newTraceId := uuid.New()

		ctx, cancel := context.WithCancel(r.Context())
		ctx = context.WithValue(ctx, "trace_id", newTraceId)

		if traceID, ok := ctx.Value(newTraceId).(string); ok {
			record.Add("trace_id", slog.StringValue(traceID))
		}

		// id := mux.Vars(newTraceId)["trace_id"]
		nextHandler.ServeHTTP(w, r)
	})

	return midHandler
}
