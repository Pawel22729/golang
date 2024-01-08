package main

import (
	"fmt"
	"time"
)

func Past(h, m, s int) int {

	h_mil := (time.Duration(h) * time.Hour).Milliseconds()
	m_mil := (time.Duration(m) * time.Minute).Milliseconds()
	s_mil := (time.Duration(s) * time.Second).Milliseconds()

	mils := h_mil + m_mil + s_mil

	return int(mils)

}

func main() {
	res := Past(1, 1, 1)
	fmt.Println(res)
}
