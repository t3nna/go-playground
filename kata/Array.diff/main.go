package main

import "fmt"

func main() {
	res := ArrayDiff([]int{1, 2, 3}, []int{})
	fmt.Println(res)
}

func ArrayDiff(a, b []int) []int {
	var res []int
outer:
	for _, num := range a {
		for i := 0; i < len(b); i++ {
			sub := b[i]
			if sub == num {
				continue outer
			}
		}
		res = append(res, num)
	}
	return res
}
