package main

import "fmt"

func main() {
	i := 5
	switch {
	case i == 5: //i=5 olduğu için diğer case’ler sorgulanmaz
		fmt.Println("i eşittir 5")
	case i < 10:
		fmt.Println("i küçüktür 10")
	case i > 3:
		fmt.Println("i büyüktür 3")
	}
}
