package main

import (
	"strings"
)

func getinitials(n string) (string, string) {
	s := strings.ToLower(n)
	names := strings.Split(s, "")
	var initials []string
	for _, v := range names {
		initials = append(initials, v[:1])
	}

	if len(initials) > 1 {
		return initials[0], initials[1]
	}
	return initials[0], "_"
}

func main() {
	fn, sn := getinitials("tifa lockhart")

	println(fn, sn)
	// Output: Ti L
}
