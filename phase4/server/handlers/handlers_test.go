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
	USERNAME1 = "ALPHA"
	USERNAME2 = "BRAVO"
	USERNAME3 = "CHARLIE"
)

func TestGetUser(t *testing.T) {
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

	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		{"Test GetUser handler: " + USERNAME1, args{httptest.NewRecorder(), httptest.NewRequest("POST", "localhost:8080/GetUser", bytes.NewBuffer(marshaledUsername1))}},
		{"Test GetUser handler: " + USERNAME2, args{httptest.NewRecorder(), httptest.NewRequest("POST", "localhost:8080/GetUser", bytes.NewBuffer(marshaledUsername2))}},
		{"Test GetUser handler: " + USERNAME3, args{httptest.NewRecorder(), httptest.NewRequest("POST", "localhost:8080/GetUser", bytes.NewBuffer(marshaledUsername3))}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GetUser(tt.args.w, tt.args.r)
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

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			GetUser(httptest.NewRecorder(), httptest.NewRequest("POST", "localhost:8080/GetUser", bytes.NewBuffer(marshaledUsername1)))
		}
	})
}
