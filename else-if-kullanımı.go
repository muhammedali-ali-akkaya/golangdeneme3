package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		if i < 5 {
			fmt.Println("sayi 5 ten küçüktür")
		} else if i == 5 {
			fmt.Println("sayi 5 e eşittir")
		} else {
			fmt.Println("sayi 5 ten büyüktür")
		}
	}
}
