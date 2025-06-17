package main

import (
	"log"
	"net/http"
)

func main() {
	// TODO: Load config
	// TODO: Init storage
	// TODO: Init scheduler
	// TODO: Init router


	log.Println("Starting RSS Reader Backend on :3010")
	http.ListenAndServe(":3010", nil)
}
