package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

func indexPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "TEST")
}

type TestData struct {
	Imie     string
	Nazwisko string
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func sqlAdd(a, b string) {
	db, err := sql.Open("sqlite3", "./baza.sql")
	checkErr(err)
	stmt, err := db.Prepare("INSERT INTO Osoby(Imie, Nazwizko) values(?,?)")
	checkErr(err)
	_, err = stmt.Exec(a, b)
	checkErr(err)
}

func doDb(query string, values []string) {
	dbDriver := "sqlite3"
	mainDb := "./baza.sql"
	db, err := sql.Open(dbDriver, mainDb)
	checkErr(err)
	stmt, err := db.Prepare(query)
	_, err = stmt.Exec()
	checkErr(err)
}

// func logInUser(username, userpass string) {

// }

func adminPage(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/admin.html"))
	if r.Method != http.MethodPost {
		tmpl.Execute(w, nil)
		return
	}
	sqlAdd(r.FormValue("imie_"), r.FormValue("nazwisko_"))
	data := TestData{Imie: r.FormValue("imie_"), Nazwisko: r.FormValue("nazwisko_")}
	tmpl.Execute(w, data)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", indexPage)
	r.HandleFunc("/admin", adminPage)
	http.ListenAndServe(":80", r)
}
