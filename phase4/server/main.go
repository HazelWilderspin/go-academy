package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"server/actor"
	"server/handlers"

	"github.com/google/uuid"
)

const (
	traceIdKey handlers.CtxKey = "trace_id"
)

func main() {
	fmt.Printf("-------------------- Go Http server starting --------------------\n")

	defaultAttrs := []slog.Attr{
		slog.String("service", "todoService"),
		slog.String("node", "auth-node-1"),
		slog.String("environment", "develop"),
	}

	slogHandler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{AddSource: true}).WithAttrs(defaultAttrs)
	slog.SetDefault(slog.New(&handlers.ContextHandler{Handler: slogHandler}))

	defer func() {
		fmt.Println("Closing RequestChannel channel")
		close(actor.RequestChannel)
	}()

	go actor.Actor()

	mux := http.NewServeMux()

	mux.Handle("/GetUser", TraceMiddleware(http.HandlerFunc(handlers.GetUser)))
	mux.Handle("/PostList", TraceMiddleware(http.HandlerFunc(handlers.PostList)))
	mux.Handle("/PutListName", TraceMiddleware(http.HandlerFunc(handlers.PutListName)))
	mux.Handle("/PutListToggleCompletion", TraceMiddleware(http.HandlerFunc(handlers.PutListToggleCompletion)))
	mux.Handle("/DeleteList", TraceMiddleware(http.HandlerFunc(handlers.DeleteList)))
	mux.Handle("/PostItem", TraceMiddleware(http.HandlerFunc(handlers.PostItem)))
	mux.Handle("/PutItem", TraceMiddleware(http.HandlerFunc(handlers.PutItem)))
	mux.Handle("/DeleteItem", TraceMiddleware(http.HandlerFunc(handlers.DeleteItem)))

	srv := &http.Server{Addr: "0.0.0.0:8080", Handler: mux}

	err := srv.ListenAndServe()
	if err != nil {
		slog.Error("Mux ListenAndServe failed", "Server Error: ", err)
	}
}

func TraceMiddleware(nextHandler http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			newTraceId := uuid.New()
			ctx := context.WithValue(r.Context(), traceIdKey, newTraceId)
			nextHandler.ServeHTTP(w, r.WithContext(ctx))
		})
}
