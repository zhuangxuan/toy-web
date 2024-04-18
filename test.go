package main

import (
	"fmt"
)

func main() {
	s := make([]int, 0, 1)
	c := cap(s)
	for i := 0; i < 1000; i++ {
		s = append(s, i)
		if n := cap(s); n > c {
			fmt.Printf("cap: %d -> %d\n", c, n)
			c = n
		}
	}
}
