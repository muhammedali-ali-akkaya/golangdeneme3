package main

import "fmt"

func main() {
	i := 5
	switch {
	case i == 5:
		fmt.Println("i eşittir 5")
		fallthrough //bu deyim ile switch in başka bir koşul daha sağlanması için kullanılır
	case i < 10:
		fmt.Println("i küçüktür 10")
	case i > 3:
		fmt.Println("i büyüktür 3")
	}
}
