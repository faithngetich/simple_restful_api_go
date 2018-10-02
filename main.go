package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
)

type Article struct {
	Title string `json:"Title"`
	Desc string `json:"desc"`
	Content string `json:"content"`
}

type Articles []Article

func AllArticles(w http.ResponseWriter, r *http.Request) {
	articles := Articles{
		Article{Title:"Test Title", Desc: "Test Description", Content: "Hello World"},
	}

	fmt.Fprintf(w, "Endpoint Hit: All Articles Endpoint")
	json.NewEncoder(w).Encode(articles)
}

func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage Endpoint Hit")
}

func HandleRequests() {
	http.HandleFunc("/", HomePage)
	http.HandleFunc("/articles", AllArticles)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

// Entry point to the application
func main() {
	HandleRequests( )
}