package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
)

func init() {
	if env := godotenv.Load(); env != nil {
		log.Panicln(fmt.Errorf("Reason: %s", env))
	}
}

func main() {
	// can be blank for now
}
