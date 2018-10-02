package main

import (
	"log"

	"github.com/vkd/goElasund/client/sdl"
)

func main() {
	log.Printf("Starting goElasund game client...")

	err := sdl.Run()
	if err != nil {
		log.Fatalf("Error on run client: %v", err)
	}
}
