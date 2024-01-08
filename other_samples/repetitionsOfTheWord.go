package main

import (
	"fmt"
	"strings"
)

func RepeatStr(repetitions int, value string) string {
	res := []string{}
	for i := 0; i <= repetitions; i++ {
		res = append(res, value)
	}
	return strings.Join(res, "")
}

func main() {
	r := RepeatStr(3, "hello ")
	fmt.Println(r)
}
