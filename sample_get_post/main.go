package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const apiUrl = "https://gorest.co.in"
const jsonHowTo = "https://www.sohamkamani.com/golang/json/"

var apiEndpoints = map[string]string{
	"users":    "/public/v2/users",
	"posts":    "/public/v2/posts",
	"comments": "/public/v2/comments",
	"todos":    "public/v2/todos",
}

var apiParams = []string{
	"per_page=100",
}

var headers = map[string]string{
	"ContentType":   "application/json",
	"Accept":        "application/json",
	"Authorization": "Bearer " + os.Getenv("AUTH_TOKEN"),
}

type User struct {
	Id     string
	Name   string
	Email  string
	Gender string
	Status string
}

func doRequest(url string, method string, headers map[string]string, data []byte) *http.Response {
	client := &http.Client{}
	client.Timeout = time.Second * 10

	reqHandler, err := http.NewRequest(method, url, nil)
	if err != nil {
		log.Fatal(err)
	}
	for key, value := range headers {
		reqHandler.Header.Add(key, value)
	}
	resp, err := client.Do(reqHandler)
	if err != nil {
		log.Fatal(err)
	}
	return resp
}

func getAllUsers() {
	lastPageReq := doRequest(apiUrl+apiEndpoints["users"], "GET", nil, nil)
	lastPage := lastPageReq.Header.Get("X-Pagination-Total")
	page := 0
	lastPageInt, _ := strconv.Atoi(lastPage)
	for page <= lastPageInt {
		additionalParams := strings.Join(apiParams, "&")
		nextPage := doRequest(apiUrl+apiEndpoints["users"]+"?page="+strconv.Itoa(page)+additionalParams, "GET", nil, nil)
		nextData, _ := io.ReadAll(nextPage.Body)
		var users []User
		json.Unmarshal([]byte(nextData), &users)
		for _, user := range users {
			fmt.Println(user.Name, ": ", user.Email)
		}
		page++
	}
}

func createUser(name string, email string, gender string, status string) {
	user := map[string]string{
		"name":   name,
		"email":  email,
		"gender": gender,
		"status": status,
	}
	jsonUser, _ := json.Marshal(user)
	resp := doRequest(apiUrl+apiEndpoints["users"], "POST", headers, []byte(jsonUser))
	fmt.Println(resp.StatusCode)
	respBody, _ := io.ReadAll(resp.Body)
	fmt.Println(string(jsonUser))
	fmt.Println(string(respBody))
}

func main() {
	//getAllUsers()
	createUser("Pawel Gonzalez", "pablo.zonzalez@gmail.com", "male", "active")
}
