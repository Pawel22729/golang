package main

import (
	"log"
	"time"
	//"encoding/json"
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

type Article struct {
	article map[string]string
	// Author    string
	// ArticleSubject string
	// ArticleContent string
	// ArticleDate string
}

func sqlInsertArticle(author string, subject string, content string, date time.Time) {
	db, err := sql.Open("sqlite3", "./baza.sql")
	checkErr(err)
	stmt, err := db.Prepare("INSERT INTO Articles(article_author,article_subject,article_content,article_date) values($1, $2, $3, $4)")
	checkErr(err)
	_, err = stmt.Exec(author, subject, content, date)
	checkErr(err)
}

func sqlGetAllArticles() []Article {
	db, err := sql.Open("sqlite3", "./baza.sql")
	checkErr(err)
	Articles, err := db.Query("SELECT * FROM Articles")
	checkErr(err)
	var ArticlesList []Article

	for Articles.Next() {
		var (
			id             int
			author         string
			articleSubject string
			articleContent string
			articleDate    string
		)
		if err := Articles.Scan(&id, &author, &articleSubject, &articleContent, &articleDate); err != nil {
			log.Fatal(err)
		}
		ArtMap := map[string]string{
			"Author": author,
			"ArticleSubject": articleSubject,
			"ArticleContent": articleContent,
			"ArticleDate": articleDate,
		}
		Art := Article{article: ArtMap} 
		ArticlesList = append(ArticlesList, Art)
	}
	return ArticlesList
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
