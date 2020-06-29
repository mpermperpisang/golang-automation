package main

import (
	"github.com/golang-automation/features/helper/errors"
	"github.com/joho/godotenv"
)

func init() {
	env := godotenv.Load()
	errors.LogPanicln(env)
}

func main() {
	// can be blank for now
}
