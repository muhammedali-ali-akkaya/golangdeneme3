package main

import "fmt"

func main() {
	metin := "merhaba ısparta"

	func(a string) {
		fmt.Println(a)
	}(metin)
}
