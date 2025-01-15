package main

import (
	"testing"
)

func TestReloadData(t *testing.T) {
	USERNAME = "BRAVO"
	name, lists := reloadData()

	if name != "Bobby" {
		t.Errorf("Username %s should have a readable name of: %s", USERNAME, name)
	}

	if len(lists) != 2 {
		t.Errorf("Username %s should have %d lists:", USERNAME, 2)
	}
}
