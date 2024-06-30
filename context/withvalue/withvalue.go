package main

import (
	context "context"
	"fmt"
)

func PrintValue(ctx context.Context) string {
	for _, key := range ctx.Value("test key").(string) {
		fmt.Println(key)
	}
	return "ok"
}

func main() {
	// create context
	baseCtx := context.Background()
	// add value
	ctxv := context.WithValue(baseCtx, "mykey", "myvalue")
	//check value
	if value, ok := ctxv.Value("mykey").(string); ok {
		fmt.Println("找到值", value)
	} else {
		fmt.Println("没有找到值")
	}
	// print context value
	fmt.Println(ctxv)

}
