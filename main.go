package main

import (
	"fmt"
	"log"
	"os"

	"github.com/bdotdub/danger-go"
	"github.com/joho/godotenv"
)

func init() {
	if env := godotenv.Load(); env != nil {
		log.Panicln(fmt.Errorf("REASON: %s", env))
	}
}

func main() {
	_, results, err := danger.Danger()
	if err != nil {
		log.Fatal(err)
	}
	defer results.Flush(os.Stdout)

	results.Message("✌️ Howdy!")
}
