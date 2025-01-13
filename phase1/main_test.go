package main

import (
	"testing"
)

func TestReloadData(t *testing.T) {
	name, data := reloadData("data4.json", "BRAVO")

	println("Name:", name)
	println("Data:", data)

}
