package main

import (
	"fmt"
	"sort"
	"strings"
	"unicode"
)

func main() {
	fmt.Println(Order("4of Fo1r pe6ople g3ood th5e the2"))

}

func Order(sentence string) string {
	parts := strings.Split(sentence, " ")

	sort.Slice(parts, func(a, b int) bool {
		return findDigit(parts[a]) < findDigit(parts[b])
	})

	return strings.Join(parts, " ")
}

func findDigit(s string) int {
	for _, r := range s {
		if unicode.IsDigit(r) {
			return int(r - '0')
		}
	}
	return 0
}
