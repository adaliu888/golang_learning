package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "Hello, 世界!"
	fmt.Println(strings.ToUpper(s))                           // HELLO, 世界!
	fmt.Println(strings.ToLower(s))                           // hello, 世界!
	fmt.Println(strings.Title(s))                             // Hello, 世界!
	fmt.Println(strings.TrimSpace(s))                         // Hello, 世界!
	fmt.Println(strings.Split(s, " "))                        // [Hello, 世界!]
	fmt.Println(strings.Join([]string{"Hello,", "世界!"}, " ")) // Hello, 世界!
	fmt.Println(strings.Contains(s, "世"))                     // true
	fmt.Println(strings.Count(s, "o"))                        // 3
	fmt.Println(strings.HasPrefix(s, "Hello"))                // true
	fmt.Println(strings.HasSuffix(s, "世界!"))                  // true
	fmt.Println(strings.Index(s, "世"))                        // 6
	fmt.Println(strings.LastIndex(s, "l"))                    // 9
	fmt.Println(strings.Repeat(s, 2))                         // Hello, 世界!Hello, 世界!
	fmt.Println(strings.Replace(s, "世", "世界", 1))             // Hello, 世界!
}
