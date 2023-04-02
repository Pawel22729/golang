package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func errChecker(err error) {
	if err != nil {
		panic(err)
	}
}

func findMaxFromFile(filePath string) int {
	file, err := os.Open(filePath)
	errChecker(err)

	defer file.Close()

	scanner := bufio.NewScanner(file)
	max := 0

	for scanner.Scan() {
		line := scanner.Text()
		num, err := strconv.Atoi(line)
		errChecker(err)
		if num > max {
			max = num
		}
	}

	return max
}

func main() {
	file := "./numbers_list.txt"
	max := findMaxFromFile(file)

	fmt.Println(max)
}
