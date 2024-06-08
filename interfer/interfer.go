package main

import "fmt"

type animal struct {
	name  string
	speak string
}

func (a *animal) Speak() {
	fmt.Println(a.speak)
}

func (a *animal) GetName() {
	fmt.Println(a.name)
}

func speak(s string) {
	fmt.Println(s)
}

func GetName(s string) string {
	return s
}

func main() {
	a := animal{
		name:  "cat",
		speak: "meow",
	}
	a.Speak()
	a.GetName()
	speak("hello")
	fmt.Println(GetName("hello"))
	speak(GetName("hello"))
	speak(GetName(GetName("hello")))
	speak(GetName(GetName(GetName("hello"))))
	speak(GetName(GetName(GetName(GetName("hello")))))
	speak(GetName(GetName(GetName(GetName(GetName("hello"))))))
	speak(GetName(GetName(GetName(GetName(GetName(GetName("hello")))))))
	speak(GetName(GetName(GetName(GetName(GetName(GetName(GetName("hello"))))))))
	speak(GetName(GetName(GetName(GetName(GetName(GetName(GetName(GetName("hello")))))))))
}
