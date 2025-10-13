package main

import (
	"log"
	"os"

	"github.com/jrkitt/ignite/ignite"
)

func main() {
	if len(os.Args) < 3 {
		log.Fatalf("Usage: ignite <path/glob> <command>\nExample: ignite './*.go' 'go run main.go'")
	}

	path := os.Args[1]
	cmd := os.Args[2]

	err := ignite.Start([]string{path}, cmd)
	if err != nil {
		log.Fatal(err)
	}
}
