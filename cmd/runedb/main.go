package main

import "github.com/rum3r/rune-db/internal/engine"

func main() {
	db := engine.New()

	db.Set("key", "value")

	val, err := db.Get("key")
	if err != nil {
		panic(err)
	}
	println(val)
}
