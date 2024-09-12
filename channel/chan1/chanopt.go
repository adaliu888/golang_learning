package chanopt

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func doWork() int {
	time.Sleep(time.Second)
	return rand.Intn(100)
}

func ChanOpt() {
	datachan := make(chan int) //make a channel ,but buffer does not infered

	go func() {
		wg := sync.WaitGroup{} // make a sync group
		for i := 0; i < 10; i++ {
			wg.Add(1) // add
			go func() {
				defer wg.Done()
				datachan <- doWork()
			}()
		}
		wg.Wait()       // wait for the channel to complete
		close(datachan) // close the channel
	}()

	for n := range datachan {
		fmt.Printf("n = %d\n", n)

	}
}
