package main

import "fmt"

func main() {
	msg := "x"

	{
		fmt.Println("1:", msg)

		msg := "y"

		fmt.Println("2:", msg)
	}

	fmt.Println("3:", msg)
}