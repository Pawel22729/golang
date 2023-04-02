package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func fetchUrl(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

type User struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Gender string `json:"gender"`
	Status string `json:"status"`
}

func unpackUsers(payload string) []User {
	var users []User
	err := json.Unmarshal([]byte(payload), &users)
	if err != nil {
		panic(err)
	}
	return users
}

func main() {
	url := "https://gorest.co.in/public/v2/users"
	resp, err := fetchUrl(url)
	if err != nil {
		fmt.Println(err)
		return
	}

	users := unpackUsers(resp)

	for _, user := range users {
		fmt.Println(user.ID)
	}
}
