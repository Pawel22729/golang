package main

import (
	"fmt"
)

func LoveFunc(flower1, flower2 int) bool {
	res := (flower1 % 2) + (flower2 % 2)
	if res == 1 {
		return false
	}
	return true
}

func main() {
	resp := LoveFunc(2, 2)
	fmt.Println(resp)
}
