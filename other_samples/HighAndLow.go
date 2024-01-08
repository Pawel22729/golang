package main

import (
	"fmt"
	"strconv"
	"strings"
)

func HighAndLow(in string) {
	lowest := 0
	highest := 0
	l := strings.Split(in, " ")
	for _, v := range l {
		num, _ := strconv.Atoi(v)
		if num > highest {
			highest = num
		}
		if num < lowest {
			lowest = num
		}
	}
	fmt.Println(strconv.Itoa(highest) + " " + strconv.Itoa(lowest))
}

func main() {
	HighAndLow("8 3 -5 42 -1 0 0 -9 4 7 4 -4")
}
