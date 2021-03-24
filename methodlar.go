package main

import "fmt"

type insan struct {
	ad  string
	yas int
}

//tanıt
func (i insan) tanıt() {
	fmt.Printf("merhaba ben %s %d yaşındayım", i.ad, i.yas)
}

func main() {
	kişi := insan{"muhammed", 24}
	kişi.tanıt()
}
