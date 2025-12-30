package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "Łukasz"
	fmt.Println(len(s))
	fmt.Println(utf8.RuneCountInString(s))

	for i := range s {
		fmt.Printf("position %d: %c\n", i, s[i])
	}
	fmt.Println("==============================")

	// We have to use the value element of the range operator
	for i, r := range s {
		fmt.Printf("position %d: %c\n", i, r)
	}
	fmt.Println("==============================")
	fmt.Println("Memory leaks with strings")

	s1 := "lorem ipsum"
	s2 := s1[:5]
	fmt.Println(s2)
}
