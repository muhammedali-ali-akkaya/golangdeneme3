package main

import "fmt"

type kişi struct {
	isim    string
	soyisim string
	yas     int
}

func main() {
	kişi1 := kişi{"muhammed", "akkaya", 24}
	kişi1.isim = "ahmet"
	kişi1.soyisim = "akbaba"
	fmt.Println(kişi1)
}
