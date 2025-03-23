package main

import (
	"fmt"
)

type Event struct {
	Name string
	NAME *Event
}

func (e *Event) Error() string {
	return fmt.Sprintf("error: %s", e.Name)
}

func main() {

	var e Event = Event{Name: "error"}
	fmt.Println(e)

}
