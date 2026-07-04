package main

import (
	"github.com/rum3r/rune-db/internal/engine"
)

func main() {
	db, err := engine.Open("runedb.log")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	if err := db.Set("key", "value"); err != nil {
		panic(err)
	}

	val, err := db.Get("key")
	if err != nil {
		panic(err)
	}
	println(val)
}
