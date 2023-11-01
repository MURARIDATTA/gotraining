package main

import (
	"fmt"
	"myproject/simplecalc"
)

func main() {
	a, b := 15.0, 5.0
	fmt.Println("hello world")
	fmt.Println(simplecalc.Add(a, b))
	fmt.Println(simplecalc.Sub(a, b))
}
