package nbuff

import (
	"fmt"
)

//create a new channel instance with no buffering

func NBuff() {
	ch := make(chan int)

	go func() {
		ch <- 1
	}()
	fmt.Println(<-ch)

}
