package main

import (
	"fmt"
	"log"
)

type Option struct {
	id    string
	value string
}

func main() {
	A := []int{1, 2, 3, 5, 8}

	for i, v := range A {
		if i == 1 {
			A[i+1] += 10
		}
		fmt.Println(v)
	}
	fmt.Println("================================")

	a := [3]int{0, 1, 2}
	log.Println(a)

	for i, v := range a {
		a[2] = 10
		if i == 2 {
			fmt.Println(v)
		}
	}

	fmt.Println("================================")

	M := map[string]Option{
		"1": {
			id:    "",
			value: "",
		},
		"2": {
			id:    "",
			value: "",
		},
	}

	M["2"] = Option{
		id:    "1",
		value: "323",
	}
	UpdateMap(M)

	fmt.Println(M)

}

func UpdateMap(m map[string]Option) {
	//m["1"].value = "wewe"

}
