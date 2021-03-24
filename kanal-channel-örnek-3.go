package main

import (
	"fmt"
	"time"
)

func main() {
	//2 adet bool deger taşıyan bir kanal olutşturalım
	k := make(chan bool, 2)

	//asenkron bir iş parçacıgı oluşturalım
	go func() {
		//5 sn beklesin
		time.Sleep(time.Second * 5)

		//k kanalına bir bool deger gönderelim
		k <- true

		//tekraradan 2 sn beklesin
		time.Sleep(time.Second * 2)

		//ve k kanalına 2. deger de göderilsin
		k <- false
	}()

	//ana iş parçacıgı k kanalına 2 deger gelene kadar bekleyecek
	fmt.Println(<-k, <-k)
	//iki bool degeride bastırmak için k kanalını 2 defa yazdırdık
}
