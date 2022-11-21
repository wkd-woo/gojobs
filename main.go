package main

import "fmt"

func main() {
	a := 2
	b := &a
	fmt.Println(a, b)
	a = 5
	fmt.Println(a, *b)
	*b = 20
	fmt.Println(a, *b)
	a = 202020
	fmt.Println(a, *b)
	fmt.Println(a, b)
	fmt.Println(&a, b)
	fmt.Println(&a, &b)
}
