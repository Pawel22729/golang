package main

import (
	"fmt"
	"strings"
)

func NoSpace(word string) string {
	newString := strings.Replace(word, " ", "", -1)
	return newString
}

func main() {
	resp := NoSpace("8 j 8   mBliB8g  imjB8B8  jl  B")
	fmt.Println(resp)
}
