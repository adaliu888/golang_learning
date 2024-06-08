package main

//this is the package that is inferred. fmt,math,and time
import (
	"fmt"
	"math"
	"time"
)

func main() {
	fmt.Println(time.Now().Local().UnixMicro())
	fmt.Println("random number generation %v", math.Abs(-1))
}
