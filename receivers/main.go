package main

import "fmt"

func main() {
	m := map[string]int{
		"1": 1,
		"2": 2,
		"3": 3,
	}

	add(&m)
	fmt.Println(m)
}

func add(a *map[string]int) {
	(*a)["b"] = 34

	fmt.Println(a)
}
