package main

import (
	"fmt"
	"time"
)

func main() {
	kanal := make(chan string) //kanal oluşturuyoruz
	go func() {
		time.Sleep(time.Second * 2) //2 sn uyku
		kanal <- "kanal bitti"      //iletişime geçiriyoruz
		fmt.Println("anonim fonsiyon yazısı")
	}()
	fmt.Println("öylesine bir yazı")
	fmt.Println(<-kanal) //kanaldan gelen veri bekleniyor
}
