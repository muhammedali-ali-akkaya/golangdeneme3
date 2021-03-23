package main

import "fmt"

func toplama(sayilar ...int) int {
	toplam := 0
	for _, n := range sayilar {
		toplam += n
	}
	return toplam
}

func main() {
	fmt.Println(toplama(5, 3, 6, 9))
}
