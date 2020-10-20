package main

import (
	"fmt"
)

type tipoDeCarlos int

var x tipoDeCarlos
var y int = 42

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)
	y = int(x)
	fmt.Println(y)
	fmt.Printf("%T", y)
}
