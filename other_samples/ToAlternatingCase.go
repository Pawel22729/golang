package main

import (
	"fmt"
	"strings"
	"unicode"
)

func ToAlternatingCase(str string) string {
	var res []string
	for _, v := range str {
		if unicode.IsUpper(v) {
			res = append(res, strings.ToLower(string(v)))
		} else {
			res = append(res, strings.ToUpper(string(v)))
		}
	}
	return strings.Join(res, "")
}

func main() {
	res := ToAlternatingCase("hello World")
	fmt.Println(res)
}
