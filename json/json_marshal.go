package main

import (
	"encoding/json"
	"fmt"
)

type js struct {
	Name    string
	Surname string
	Age     int
}

func main() {
	m := js{"Pablo", "Quepasa", 123}
	b, _ := json.Marshal(m)

	fmt.Println(string(b))
}
