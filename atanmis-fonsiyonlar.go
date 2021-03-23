package main

import "fmt"

func main() {
	toplam := func(a, b int) int {
		return a + b
	}
	fmt.Println(toplam(10, 30))
}
