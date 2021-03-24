package main

import "fmt"

func main() {
	kişi := struct {
		ad, soyad string
	}{"muhammed", "akkaya"}

	fmt.Println(kişi.ad, kişi.soyad)
}
