package main

import "fmt"

func main() {
	dizi := [5]int{1, 2, 3, 4, 5}
	var dizia []int
	fmt.Println(dizi)

	var dizi2 []int = dizi[2:4]
	var dizi3 []int = dizi[:3]
	var dizi4 []int = dizi[3:]
	fmt.Println(dizi2)
	fmt.Println(dizi3)
	fmt.Println(dizi4)

	fmt.Println("dizinin uzunluğu", len(dizi))
	fmt.Println("dizinin geriye kalan kapasitesi", cap(dizi))
	fmt.Println("dizinin elemanları", dizi)

	fmt.Println("dizinin uzunlugu", len(dizi2))
	fmt.Println("dizinin geriye kalan kapasitesi", cap(dizi2))
	fmt.Println("dizinin elemanları", (dizi2))

	if dizia == nil {
		fmt.Println("dizia boş dizi")
	}
	a := make([]int, 5)    //uzunluğu 5 olan dizi
	b := make([]int, 0, 5) //uzunluğu 0 kapasitesi 5 olan dizi
	fmt.Println(a)
	fmt.Println(b)
	//diziye ekleme yapma işlemleri
	var dizib []string
	dizib = append(dizib, "eklenen eleman 1")
	dizib = append(dizib, "eklenen eleman 2")
	fmt.Println(dizib)
	fmt.Println("dizib uzunluk ve kapsam", len(dizib), cap(dizib))

}
