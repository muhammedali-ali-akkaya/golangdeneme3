package main

import "fmt"

func main() {
	m := make(map[string]int) //m isminde string bölge isimli int deger taşıyan dizi
	m["sayi"] = 25            //
	fmt.Println(m)
	fmt.Println(m["sayi"])
	delete(m, "sayi") //sayi bölgesindeki degeri sildik
	fmt.Println(m["sayi"])
}
