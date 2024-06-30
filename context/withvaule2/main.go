package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

func main() {
	start := time.Now()
	//ctx := context.Background()
	//withValues
	ctx := context.WithValue(context.Background(), "foo", "bar") //transfort key and value ,withValue means that is
	userID := 1
	val, err := fetchUserData(ctx, userID)
	if err != nil { //check error
		log.Fatal(err)
	}
	fmt.Println("result", val)
	fmt.Println("took", time.Since(start))
}

type Respones struct {
	value int
	err   error
}

// context.context is used to share information between goroutines
func fetchUserData(ctx context.Context, userID int) (int, error) { //
	val := ctx.Value("foo")
	fmt.Println(val)
	//创建一个带有超时时间的context，如果超时时间的context，如果超时时间到达，这个context就会被取消
	ctx, cancel := context.WithTimeout(ctx, time.Millisecond*200)
	defer cancel()
	if ctx.Err() != nil { //check error
		return 0, ctx.Err()
	}
	respch := make(chan Respones) //make channel,and do something,channel send information
	go func() {
		val, err := fetchThirdPartystuffWhichCanBeSlow()
		respch <- Respones{value: val, err: err}
	}()
	for {
		select {

		case <-ctx.Done(): //返回一个channel，当context被取消时，这个channel会被关闭
			return 0, fmt.Errorf("fetching data from third party took too long")
		case resp := <-respch:
			return resp.value, resp.err

		}
	}
}

// fetch time of data from third party ,here is set 150ms
func fetchThirdPartystuffWhichCanBeSlow() (int, error) {
	time.Sleep(time.Millisecond * 150)
	return 666, nil
}
