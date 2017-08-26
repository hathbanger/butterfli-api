package main

import (
	"github.com/hathbanger/butterfli-api/server"
)

func main() {
	// TODO: Refector all error handling.. way to much panicing
	// TODO: Implement logging
	// TODO: Implement PostIds in account struct instead of posts composition.. will need to seperate Posts into new collection.. ect
	// TODO: Clone store instead of creating new handler.. not DRY
	// TODO: Refector echo implementation to use most recent versioning.. code has regressed out of date

	server.Run()
}