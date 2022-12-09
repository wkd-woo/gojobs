package main

import "fmt"

func main() {
	jae := map[string]string{"name": "jae", "age": "12"}
	for _, value := range jae {
		fmt.Println(value)
	}
}
