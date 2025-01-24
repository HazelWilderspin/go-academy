package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"server/actor"
	"server/crud"
	"testing"
)

var (
	USERNAME1    = "ALPHA"
	USERNAME2    = "BRAVO"
	USERNAME3    = "CHARLIE"
	UNKNOWN_USER = "UNKNOWN_USER"
)

func trace(name string) func() {
	fmt.Printf("%s --------- entered", name)
	return func() {
		fmt.Printf("%s --------- returned", name)
	}
}

func TestGetUser(t *testing.T) {
	defer trace("TestReadItem")()
	t.Parallel()

	defer func() {
		fmt.Println("Closing RequestChannel channel")
		close(actor.RequestChannel)
	}()

	go actor.Actor()

	crud.FILE_PATH = "../crud/mockData.json"

	marshaledUsername1, err := json.Marshal(GetUserRequestBody{Username: USERNAME1})
	if err != nil {
		log.Fatal(err)
	}
	marshaledUsername2, err := json.Marshal(GetUserRequestBody{Username: USERNAME2})
	if err != nil {
		log.Fatal(err)
	}
	marshaledUsername3, err := json.Marshal(GetUserRequestBody{Username: USERNAME3})
	if err != nil {
		log.Fatal(err)
	}
	marshaledUnknownUser, err := json.Marshal(GetUserRequestBody{Username: UNKNOWN_USER})
	if err != nil {
		log.Fatal(err)
	}

	type args struct {
		r *http.Request
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"Test GetUser handler: " + USERNAME1, args{httptest.NewRequest("GET", "localhost:8080/GetUser", bytes.NewBuffer(marshaledUsername1))}, false},
		{"Test GetUser handler: " + USERNAME2, args{httptest.NewRequest("GET", "localhost:8080/GetUser", bytes.NewBuffer(marshaledUsername2))}, false},
		{"Test GetUser handler: " + USERNAME3, args{httptest.NewRequest("GET", "localhost:8080/GetUser", bytes.NewBuffer(marshaledUsername3))}, false},
		{"Test GetUser handler: " + UNKNOWN_USER, args{httptest.NewRequest("GET", "localhost:8080/GetUser", bytes.NewBuffer(marshaledUnknownUser))}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			GetUser(w, tt.args.r)
			if (w.Code != http.StatusOK) != tt.wantErr {
				t.Errorf("ReadItem() Writer code = %v, StatusOK = %v, wantErr %v", w.Code, http.StatusOK, tt.wantErr)
				return
			}
		})
	}
}

func BenchmarkGetUser(b *testing.B) {
	go actor.Actor()

	crud.FILE_PATH = "../crud/mockData.json"

	marshaledUsername1, err := json.Marshal(GetUserRequestBody{Username: USERNAME1})
	if err != nil {
		log.Fatal(err)
	}
	marshaledUnknownUser, err := json.Marshal(GetUserRequestBody{Username: UNKNOWN_USER})
	if err != nil {
		log.Fatal(err)
	}

	tests := []struct {
		name    string
		user    []byte
		wantErr bool
	}{
		{"Benchmark GetUser handler: " + USERNAME1, marshaledUsername1, false},
		{"Benchmark GetUser handler: " + UNKNOWN_USER, marshaledUnknownUser, true},
	}
	for _, tt := range tests {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				w := httptest.NewRecorder()
				GetUser(w, httptest.NewRequest("GET", "localhost:8080/GetUser", bytes.NewBuffer(tt.user)))
				if (w.Code != http.StatusOK) != tt.wantErr {
					b.Errorf("ReadItem() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
			}
		})
	}
}
