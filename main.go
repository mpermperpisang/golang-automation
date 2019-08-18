package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/logrusorgru/aurora"
)

func init() {
	if env := godotenv.Load(); env != nil {
		log.Fatalln(aurora.Bold(aurora.Red(env)))
	}
}

func main() {
	// can be blank for now
}
