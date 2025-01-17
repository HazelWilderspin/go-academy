package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"

	h "phase2/handlers"

	"github.com/google/uuid"
)

const (
	traceIdKey h.CtxKey = "trace_id"
)

func main() {
	fmt.Printf("-------------------- Go Http server starting --------------------\n")

	defaultAttrs := []slog.Attr{
		slog.String("service", "todoService"),
		slog.String("node", "auth-node-1"),
		slog.String("environment", "develop"),
	}

	slogHandler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{AddSource: true}).WithAttrs(defaultAttrs)
	slog.SetDefault(slog.New(&h.ContextHandler{Handler: slogHandler}))

	mux := http.NewServeMux()

	mux.Handle("/GetUser", TraceMiddleware(http.HandlerFunc(h.GetUser)))
	mux.Handle("/GetList", TraceMiddleware(http.HandlerFunc(h.GetList)))
	mux.Handle("/PostList", TraceMiddleware(http.HandlerFunc(h.PostList)))
	mux.Handle("/PutListName", TraceMiddleware(http.HandlerFunc(h.PutListName)))
	mux.Handle("/PutListToggleCompletion", TraceMiddleware(http.HandlerFunc(h.PutListToggleCompletion)))
	mux.Handle("/DeleteList", TraceMiddleware(http.HandlerFunc(h.DeleteList)))
	mux.Handle("/GetItem", TraceMiddleware(http.HandlerFunc(h.GetItem)))
	mux.Handle("/PostItem", TraceMiddleware(http.HandlerFunc(h.PostItem)))
	mux.Handle("/PutItem", TraceMiddleware(http.HandlerFunc(h.PutItem)))
	mux.Handle("/DeleteItem", TraceMiddleware(http.HandlerFunc(h.DeleteItem)))

	srv := &http.Server{Addr: "0.0.0.0:8080", Handler: mux}

	err := srv.ListenAndServe()
	if err != nil {
		slog.Error("Bufio reader failed", "Server Error: ", err)
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
