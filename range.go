package main

import "fmt"

var isimler = []string{"ali", "veli", "hasan", "enes", "abdulkadir", "omer"}

func main() {

	for a, b := range isimler {
		fmt.Println("%d %s", a, b)
	}
}
