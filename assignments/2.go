package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter first name:")
	forename, err := reader.ReadString('\n')
	reportIfError(err)

	fmt.Println("Enter middle name:")
	middleName, err := reader.ReadString('\n')
	reportIfError(err)

	fmt.Println("Enter surname:")
	surname, err := reader.ReadString('\n')
	reportIfError(err)

	fmt.Printf("Name provided: %s %s %s", forename, middleName, surname)
}

func reportIfError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
