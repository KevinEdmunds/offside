package main

import (
	"fmt"
	"log"
	"offside/internal/store"
)

func main() {
    db, err := store.Connect()
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    fmt.Println("Connected to database successfully")
}