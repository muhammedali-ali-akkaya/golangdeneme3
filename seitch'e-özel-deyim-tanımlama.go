package main

import "fmt"

func main() {
	switch a := 5; {
	case a < 4:
		fmt.Println("a 4 e eşittir")
	case a == 5:
		fmt.Println("a 5 e eşittir")
	}
}
