package main

import "fmt"

func main() {
	defer fmt.Println("birinci cümle")
	defer fmt.Println("ikici cümle")
	defer fmt.Println("ücüncü  cümle")
	defer fmt.Println("dördüncü cümle")
	fmt.Println("besinci cümle")
}
