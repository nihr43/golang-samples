package main

import "fmt"

type anyslice[T any] []T

func (s anyslice[T]) first() T {
	return s[0]
}

func main() {
	s1 := anyslice[int]{1, 2, 3}
	s2 := anyslice[string]{"a", "b", "c"}

	fmt.Println(s1.first())
	fmt.Println(s2.first())
}
