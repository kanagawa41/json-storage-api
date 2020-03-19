package main

import (
	"fmt"

	scribble "github.com/nanobox-io/golang-scribble"
)

func test() {
	db, err := scribble.New("dir", nil)
	if err != nil {
		fmt.Println("Error", err)
	}

	// Delete all fish from the database
	if err := db.Delete("fish", ""); err != nil {
		fmt.Println("Error", err)
	}
}
