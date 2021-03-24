package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		time.Sleep(time.Second * 2)
		fmt.Println("ilk yazımız")
	}() // ( yerleştirmek zorundayız)

	fmt.Println("ikinci yazımız")
	time.Sleep(time.Second * 3)
}
