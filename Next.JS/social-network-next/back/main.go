package main

import (
	"social/backend"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	backend.StartServer()
}
