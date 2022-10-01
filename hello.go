package main

import (
	"fmt"
	"math"
)

func main() {
	s := make([]string, 3)

	s = append(s, "rohan", "utsav", "subu")

	fmt.Println(s)

	m := make(map[string]int)

	m["rohan"] = 2
	fmt.Println(m)
	// fmt.Println("emp:", s)
	// s[0] = "a"
	// s[1] = "b"
	// s[2] = "c"
	// fmt.Println("set:", s)
	// fmt.Println("get:", s[2])
	// s = append(s, "d")
	// s = append(s, "e", "f")
	// fmt.Println("apd:", s)
	// c := make([]string, len(s))
	// copy(c, s)
	// fmt.Println("cpy:", c)

	math.Abs(1.5)

	fmt.Println(math.Abs(1.5))
}
