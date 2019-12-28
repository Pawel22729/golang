package main

import (
	"html/template"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func indexPage(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	var ArticlesData = sqlGetAllArticles()
	tmpl.Execute(w, ArticlesData)
}

func addArticle(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/admin.html"))
	if r.Method != http.MethodPost {
		tmpl.Execute(w, nil)
		return
	}
	dateNow := time.Now()
	if r.FormValue("form_author") != "" && r.FormValue("form_subject") != "" && r.FormValue("form_content") != "" {
		sqlInsertArticle(r.FormValue("form_author"), r.FormValue("form_subject"), r.FormValue("form_content"), dateNow)
		var TestData = make(map[string]string)
		TestData["Author"] = r.FormValue("form_author")
		tmpl.Execute(w, TestData)
	}
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", indexPage)
	r.HandleFunc("/add", addArticle)
	http.ListenAndServe(":8080", r)
}
