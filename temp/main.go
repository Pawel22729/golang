package main

import (
	"strings"
	"fmt"
)

func Solution(word string) string {

  var new_word []string
  l := len(word) - 1
  
  for i := l; i >= 0; i-- {
    new_word = append(new_word, string(word[i]))
  }
  return strings.Join(new_word, "")
}

func main() {
	s := Solution("DUPA")
	fmt.Println(s)
}
