package main

import (
	"fmt"
)

func First(f string) (string, error) {
	if f == "" {
		return "", fmt.Errorf("First input string is empty")
	}
	return f, nil

}

func Second(s string) (string, error) {
	if s == "" {
		return "", fmt.Errorf("second input string is empty")
	}
	return s, nil
}

func Third(t string) (string, error) {
	r, err := Second(t)
	if err != nil {
		return "", err
	}
	return r, nil
}

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Panic:", err)
		}
	}()
	r, err := First("")
	if err != nil {
		panic("First input string is empty")
	}
	s, err := Second(r)
	if err != nil {
		panic(err)
	}
	fmt.Println(s)

}
