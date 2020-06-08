package main

import (
	"fmt"
	"html/template"
	"net/http"
        
	"github.com/gorilla/mux"
)

func testHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	title := vars["title"]
	page := vars["page"]

	fmt.Fprintf(w, "Title %s, page %s\n", title, page)
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
    stock := get_stock("BAC")
    fmt.Println(stock)
	tmpl, _ := template.ParseFiles("templates/index.html")
	tmpl.Execute(w, stock)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", IndexHandler).Methods("GET")
	r.PathPrefix("/files/").Handler(http.StripPrefix("/files/", http.FileServer(http.Dir("./files/"))))
	http.ListenAndServe(":8000", r)
}
