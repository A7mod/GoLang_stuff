package main

import (
	"log"

	"github.com/a7mod/server/db"
)

func main() {
	_, err := db.NewDatabase()
	if err != nil {
		log.Fatalf("couldn't initialize database connection : %s", err)
	}
}
