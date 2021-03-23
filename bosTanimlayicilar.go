package main

import "fmt"

func fonksiyon(girdi int) (int, int) {
	islem1 := girdi / 2
	islem2 := girdi / 4
	return islem1, islem2
}

func main() {
	ikiyeböl, _ := fonksiyon(16)
	fmt.Println(ikiyeböl)
}
