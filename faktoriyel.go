package main

import "fmt"

func main() {
	fmt.Println(faktoriyel(5))
}
func faktoriyel(a uint) uint {
	if a == 0 {
		return 1
	}
	return a * faktoriyel(a-1)
}
