package main

import "fmt"

func main() {
	for i := 0; i < 20; i++ {
		switch i {
		case 2:
			fmt.Println("sayi 2 den geçti")
		case 5:
			fmt.Println("sayi 5 ten geçti")
		case 9:
			fmt.Println("sayi 10 a geldi ve durdu")
			//default:
			//	fmt.Println("varsayılan deger çalıştı")

		}
	}
}
