package main

import (
	"fmt"
	"strconv"
)

func main() {
	v := "42"
	fmt.Printf("%T, %s", v, v)
	if s, err := strconv.ParseUint(v, 10, 32); err == nil {
		fmt.Printf("%T, %v\n", s, s)
	}
	if s, err := strconv.ParseUint(v, 10, 64); err == nil {
		fmt.Printf("%T, %v\n", s, s)
	}

}
