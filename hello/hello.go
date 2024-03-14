package hello

import "fmt"

func SayHello() {
	fmt.Println("Hello")
}

func SayGoodbye() {
	fmt.Println("Goodbye")
}

func SayHelloTo(name string) {
	fmt.Println("Hello", name)
}
