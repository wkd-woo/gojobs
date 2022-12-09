package main

import "fmt"

type person struct {
	name    string
	age     int
	favFood []string
}

func main() {
	favfood := []string{"beef", "coffee"}
	jae := person{name: "jae", age: 25, favFood: favfood}
	fmt.Println(jae)
}
