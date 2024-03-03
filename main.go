package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("hello world")

	// Need to run `export PORT=8000`
	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("PORT is not found in the environment")
	}

	fmt.Println("Port:", portString)
}
