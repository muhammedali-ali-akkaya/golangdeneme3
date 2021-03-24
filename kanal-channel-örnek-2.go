package main

import "time"

func main() {
	//bir kanal oluşturalım
	k := make(chan bool)
	//bu kanalımız bool deger taşıyacak

	//asenkron bir iş parçacıgı oluşturalım
	go func() {
		time.Sleep(time.Second * 5) //bu iş parçacıgı 5 sn beklesin

		k <- true
	}()

	//ana iş parçacıgı k kanalınadeger gelene kadar bekleyecek
	<-k
	//deger geldiginde program sonlanacak
}
