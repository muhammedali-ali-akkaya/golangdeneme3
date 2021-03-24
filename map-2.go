package main

import "fmt"

type insan struct {
	kişi1, kişi2, kişi3 string
}

var m = map[string]insan{
	"erkekler": {"ali", "veli", "ahmet"},
	"bayanlar": {"hatice", "zehra", "betül"},
}

func main() {
	fmt.Println(m["erkekler"])
	fmt.Println(m["bayanlar"])
	fmt.Println(m)
}
