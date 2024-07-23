package main

import "log"

func main() {
	//select {} // all goroutines are asleep - deadlock!
	log.Println("hello")

}
