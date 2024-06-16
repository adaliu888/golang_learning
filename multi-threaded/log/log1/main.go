package main

import (
	"errors"
	"log"
	"os"
)

func someFunction() error {
	return errors.New("some error")
}
func main() {
	log.SetOutput(os.Stdout)
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	log.Println("Hello, World!")
	if errors := someFunction(); errors != nil {
		log.Fatal("Fatal error: ", errors)
	}
}
