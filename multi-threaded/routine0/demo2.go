package main

//
import (
	"log"
	"math/rand"
	"sync"
	"time"
)

func SayGreetings(greeting string, times int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < times; i++ {
		log.Printf("times: %d\n", i)
		log.Println(greeting)
		d := time.Second * time.Duration(rand.Intn(5))
		log.Printf("sleep: %s\n", d) //
		time.Sleep(d)                // 睡眠片刻（随机0到2.5秒）
	}
}
func main() {
	//goutine add counter
	var wg sync.WaitGroup
	wg.Add(2)
	rand.Seed(time.Now().UnixNano()) // Go 1.20之前需要
	log.SetFlags(0)
	go SayGreetings("hi!", 10, &wg)
	go SayGreetings("hello!", 10, &wg)
	wg.Wait() // 等待所有goroutine完成。
	time.Sleep(2 * time.Second)
}
