package main

import (
	"github.com/joho/godotenv"
	"github.com/mpermperpisang/golang-automation-v1/features/helper"
)

func init() {
	env := godotenv.Load()
	helper.LogPanicln(env)
}

func main() {
	// can be blank for now
}
