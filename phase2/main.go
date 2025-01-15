package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"

	h "phase2/handlers"
)

func main() {
	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{AddSource: true})))
	fmt.Println("Go Http server starting")

	mux := http.NewServeMux()

	// temp: load json into mem store to test api
	// mux.Handle("/LoadData", middleware(h.LoadData))

	finalHandler := http.HandlerFunc(h.LoadData)
	mux.Handle("/LoadData", middleware(finalHandler))

	mux.HandleFunc("/ListHeaders", h.ListHeaders)

	mux.HandleFunc("/GetUser", h.GetUser)
	// mux.HandleFunc("/GetLists", h.GetLists)
	// mux.HandleFunc("/GetList", h.GetList)

	srv := &http.Server{Addr: "0.0.0.0:8080", Handler: mux}

	err := srv.ListenAndServe()
	if err != nil {
		slog.Error("Bufio reader failed", "Server Error: ", err)
	}
}

func middleware(nexHandler http.Handler) http.Handler {

	midHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// gen a trace id to use in ctx
		nexHandler.ServeHTTP(w, r)
	})

	return midHandler
}
