package main

import "fmt"

func main() {
	a := 8
	ekle(5)
	fmt.Println(a)
	fmt.Println(&a)
	b := &a
	fmt.Println(*b)
}
func ekle(a int) {
	a = a + 5
}
