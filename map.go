package main

import "fmt"

type insan struct {
	kişi1, kişi2, kişi3 string
}

func main() {
	var m map[string]insan
	m = make(map[string]insan)
	m["isim"] = insan{"ali", "ahmet", "mehmet"}
	fmt.Print(m["isim"])
}
