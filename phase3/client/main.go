package main

import (
	"embed"
	"fmt"
	"log/slog"
	"net/http"
	"os"

	h "client/handlers"
)

var (
	//go:embed static
	static embed.FS
)

type ContextHandler struct {
	slog.Handler
}

func main() {
	fmt.Printf("-------------------- Go Http client starting --------------------\n")

	defaultAttrs := []slog.Attr{
		slog.String("client", "todoClient"),
		slog.String("node", "auth-node-2"),
		slog.String("environment", "develop"),
	}

	slogHandler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{AddSource: true}).WithAttrs(defaultAttrs)
	slog.SetDefault(slog.New(&ContextHandler{Handler: slogHandler}))

	mux := http.NewServeMux()

	mux.Handle("/static/", http.FileServer(http.FS(static)))

	mux.Handle("/login", http.HandlerFunc(h.HomePageHandler))
	mux.Handle("/myLists", http.HandlerFunc(h.LoginHandler))
	mux.Handle("/addListForm", http.HandlerFunc(h.NewListHandler))
	mux.Handle("/submitNewList", http.HandlerFunc(h.SubmitListFormHandler))

	srv := &http.Server{Addr: "0.0.0.0:8090", Handler: mux}

	err := srv.ListenAndServe()
	if err != nil {
		slog.Error("Mux ListenAndServe failed", "Client Error: ", err)
	}
}
