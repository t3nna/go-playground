package main

import "fmt"

func main() {
	s1 := make([]int, 3, 6)
	s1[0] = 0
	s1[1] = 1
	s1[2] = 2

	s2 := s1[1:3]

	fmt.Println(s1)
	fmt.Println(s2)

	s2[0] = 10

	fmt.Println(s1)
	fmt.Println(s2)

	fmt.Println("====================")

	sl1 := make([]int, 3, 6)
	sl1[0] = 1
	sl1[1] = 2
	sl1[2] = 3
	sl2 := sl1

	fmt.Printf("%p \n", sl1)
	fmt.Printf("%p \n", sl2)

	sl2 = append(sl2, 4)

	fmt.Println("====================")

	fmt.Printf("%p \n", sl1)
	fmt.Printf("%p \n", sl2)

	fmt.Println("====================")
	slicesLenAndCap()
}

func slicesLenAndCap() {
	s1 := make([]int, 3, 6)
	s2 := []int{1, 2, 3}

	fmt.Printf("%v, cap - %v, len - %v \n", s1, cap(s1), len(s1))
	fmt.Printf("%v, cap - %v, len - %v \n", s2, cap(s2), len(s2))

	s3 := s1[1:]

	fmt.Printf("s1 %v, cap - %v, len - %v \n", s1, cap(s1), len(s1))
	fmt.Printf("s3 %v, cap - %v, len - %v \n", s3, cap(s3), len(s3))

	s1[1] = 55

	fmt.Printf("s1 %v, cap - %v, len - %v \n", s1, cap(s1), len(s1))
	fmt.Printf("s3 %v, cap - %v, len - %v \n", s3, cap(s3), len(s3))

	s1[1] = 0
	s1 = append(s1, 1, 2)
	s3 = append(s3, 11, 22)
	fmt.Printf("s1 %v, cap - %v, len - %v \n", s1, cap(s1), len(s1))
	fmt.Printf("s3 %v, cap - %v, len - %v \n", s3, cap(s3), len(s3))

}
