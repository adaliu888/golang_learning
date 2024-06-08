package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	timeoutContext, cancel := context.WithTimeout(context.Background(), time.Millisecond)
	defer cancel()
	//create Http request
	req, err := http.NewRequestWithContext(timeoutContext, http.MethodGet, "http://www.baidu.com", nil)
	if err != nil {
		panic(err)
	}
	//perform http request, use default client
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	//defer close connection
	defer res.Body.Close()
	//get data from http response
	imageData, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	fmt.Printf("download image of size %d\n", len(imageData))
}
