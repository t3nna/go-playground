package main

import (
	"fmt"
	"maps"
)

func main() {
	m := map[int]bool{
		0: true,
		1: false,
		2: true}
	// Produces random result
	// If a map entry is created during iteration, it may be produced during the iteration or skipped.
	// The choice may vary for each entry created and from one iteration to the next.

	for k, v := range m {
		if v {
			m[10+k] = true
		}
	}
	fmt.Println(m)
	fmt.Println("==============================")
	mapPredictable()
}

func mapPredictable() {
	m := map[int]bool{
		0: true,
		1: false,
		2: true,
	}
	m2 := maps.Clone(m)
	for k, v := range m {
		m2[k] = v
		if v {
			m2[10+k] = true
		}
	}
	fmt.Println(m2)
}
