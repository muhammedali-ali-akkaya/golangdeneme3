package main

import "fmt"

func main() {
	fmt.Println("sayıyor")
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("bitti")
}
