package main

import "fmt"

type kişi struct {
	ad    string
	soyad string
	yas   int
}

func main() {
	kişi1 := kişi{}
	kişi1.ad, kişi1.soyad = "muhammed ali", "akkaya"
	kişi1.yas = 24

	kişi2 := kişi{ad: "ahmet", soyad: "said", yas: 25}

	fmt.Println(kişi2.ad, kişi2.soyad, kişi2.yas)
	fmt.Println(kişi1.ad, kişi1.soyad, kişi1.yas)

}
