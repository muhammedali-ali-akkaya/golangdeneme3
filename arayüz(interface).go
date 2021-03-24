package main

import "fmt"

type hesap interface {
	hesapla()
}
type toplam struct {
	sayi1 int
	sayi2 int
}
type carpim struct {
	sayi1 int
	sayi2 int
}

func (t *toplam) hesapla() {
	fmt.Println(t.sayi1 + t.sayi2)
}
func (c *carpim) hesapla() {
	fmt.Println(c.sayi1 * c.sayi2)
}
func hesapyap(h hesap) {
	h.hesapla()
}

func main() {
	islem1 := toplam{5, 10}
	islem2 := carpim{5, 10}
	var işlem hesap
	işlem = &islem1

	hesapyap((işlem))

	işlem = &islem2
	hesapyap((işlem))
}
