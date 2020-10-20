package main

import (
	"fmt"
)

var y int = 42
var z string = "shaken, not stirred"

func main() {
	fmt.Println(y)
	fmt.Printf("%T\n", y)
	fmt.Println(z)
	fmt.Printf("%T\n", z)
	z = "aoeu"
}
