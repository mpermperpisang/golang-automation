package main

import (
	"github.com/golang-automation/features/helper"
	"github.com/joho/godotenv"
)

func init() {
	env := godotenv.Load()
	helper.LogPanicln(env)
}

func main() {
	// can be blank for now
	// only for testing
}
