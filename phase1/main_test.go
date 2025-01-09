package main

import (
	"testing"
)

func TestGetUserData(t *testing.T) {
	response, err := getUserData("ALPHA")
	if err != nil {
		t.Error("HAZZARD")
	} else {
		println(response)
	}

}
