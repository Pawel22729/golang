package main

import (
	"encoding/json"
	"fmt"
)

type js struct {
	Name    string
	Surname string
	Age     int
	Items   []string
	Dict    map[string]string
}

type js2 struct {
	Parameter map[string]string `json:"Parameter"`
}

type js3 struct {
	Name string `json:"name"`
}

func main() {

	j := `{
    "Name": "Pablo",
    "Surname": "Quepasa",
    "Age": 30,
    "Items": ["item1", "item2", "item3"],
    "Dict": {
      "dict1": "value1"
    }
  }`

	i := `{
    "Parameter": {
      "ARN": "arn:aws:ssm:eu-west-2:12345678:parameter/pablo/test/param",
      "LastModifiedDate": "2019-03-11 13:04:27 +0000 UTC",
      "Name": "/pablo/test/param",
      "Type": "String",
      "Value": "val123",
      "Version": "1"
    }
  }`

	k := `{
    "name": "windows"
  }`

	res := js{}
	err := json.Unmarshal([]byte(j), &res)
	if err != nil {
		panic(err)
	}

	fmt.Println(res)

	res2 := js2{}
	err = json.Unmarshal([]byte(i), &res2)
	if err != nil {
		panic(err)
	}

	fmt.Println(res2)

	res3 := js3{}
	err = json.Unmarshal([]byte(k), &res3)
	if err != nil {
		panic(err)
	}

	fmt.Println(res3)
}
