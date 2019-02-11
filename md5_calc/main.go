package main

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"net/http"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func fetchFile(url string) []byte {
	resp, err := http.Get(url)
	check(err)

	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)
	check(err)
	return content
}

func saveContent(fileName string, content []byte) {
	err := ioutil.WriteFile("/tmp/"+fileName, content, 0644)
	check(err)
}

func checkMd5(filePath string) string {
	file, err := ioutil.ReadFile(filePath)
	check(err)
	md5sum := md5.Sum(file)
	s := string(md5sum[:])
	return s
}

func main() {
	content := fetchFile("https://ifconfig.me/ip")
	saveContent("pablo_file.txt", content)
	md5 := checkMd5("/tmp/pablo_file.txt")
	fmt.Printf("%x", md5)
}
