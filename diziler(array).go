package main

import "fmt"

func main() {
	var dizi [5]string
	for i := 0; i < 5; i++ {
		if i < 2 {
			dizi[i] = "muhammed başta"
		} else if i < 4 {
			dizi[i] = "muhammed ortada"
		} else {
			dizi[i] = "muhammed sonda"
		}
	}
	for i := 0; i < 5; i++ {
		fmt.Println(dizi[i])
	}
	dizi1 := [3]string{"muhammed", "ali", "akkaya"}
	for i := 0; i < 3; i++ {
		fmt.Println(dizi1[i])
	}

}
