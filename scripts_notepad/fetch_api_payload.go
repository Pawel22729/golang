package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

var apiConf = map[string]string{
	"baseUrl": "https://gorest.co.in/public/v2",
	"users":   "/users",
}

func errChecker(err error) {
	if err != nil {
		panic(err)
	}
}

func fetchUrl(url string) (string, error) {
	resp, err := http.Get(url)
	errChecker(err)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	errChecker(err)

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
	errChecker(err)
	return users
}

func unpackUser(payload string) User {
	var user User
	err := json.Unmarshal([]byte(payload), &user)
	errChecker(err)
	return user
}

func getUser(userId int) User {
	url := apiConf["baseUrl"] + apiConf["users"] + "/" + strconv.Itoa(userId)
	resp, err := fetchUrl(url)
	errChecker(err)
	user := unpackUser(resp)
	return user
}

func main() {
	url := apiConf["baseUrl"] + apiConf["users"]
	resp, err := fetchUrl(url)
	errChecker(err)

	users := unpackUsers(resp)

	for _, user := range users {
		fmt.Println("User ID: ", user.ID)
		userData := getUser(user.ID)

		fmt.Println(userData)
	}
}
