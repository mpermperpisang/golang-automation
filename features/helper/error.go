package helper

import (
	"fmt"
	"log"
)

func LogFatal(err interface{}) error {
	if err != nil {
		log.Fatal(fmt.Errorf("REASON: %s", err))
	}

	return nil
}

func LogFatalln(err interface{}) error {
	if err != nil {
		log.Fatalln(fmt.Errorf("REASON: %s", err))
	}

	return nil
}

func LogPanic(err interface{}) error {
	if err != nil {
		log.Panic(fmt.Errorf("REASON: %s", err))
	}

	return nil
}

func LogPanicln(err interface{}) error {
	if err != nil {
		log.Panicln(fmt.Errorf("REASON: %s", err))
	}

	return nil
}

func LogPrint(err interface{}) error {
	if err != nil {
		log.Print(fmt.Errorf("REASON: %s", err))
	}

	return nil
}

func Println(err interface{}) error {
	if err != nil {
		log.Println(fmt.Errorf("REASON: %s", err))
	}

	return nil
}
