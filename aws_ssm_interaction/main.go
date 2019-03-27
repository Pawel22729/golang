package main

import (
	"flag"
	"fmt"
)

func main() {
	del := flag.Bool("del", false, "delete flag - bool")
	put := flag.Bool("put", false, "insert parameter - bool")
	get := flag.Bool("get", false, "get single param")
	gets := flag.Bool("gets", false, "get params by path")
	param := flag.String("param", "", "parameter path")
	val := flag.String("value", "", "parameter value")

	flag.Parse()

	if *del {
		res := delParam(*param)
		fmt.Println(res)
	}

	if *put {
		res := putParam(*param, *val)
		fmt.Println(res)
	}

	if *get {
		res := getParam(*param)
		fmt.Println(res)
	}

	if *gets {
		res := getParams(*param)
		fmt.Println(res)
	}
}
