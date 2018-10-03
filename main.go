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

// A ResponseWriter interface is used by an HTTP handler to construct an HTTP response
// A Request represents an HTTP request received by a server or to be sent by a client.
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

// HandleFunc registers the handler function for the given pattern in the DefaultServeMux
// ListenAndServe listens on the TCP network address addr and then calls Serve with handler to handle requests on incoming connection. It's starts an HTTP server with a given address and handler
func HandleRequests() {
	http.HandleFunc("/", HomePage)
	http.HandleFunc("/articles", AllArticles)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

// Entry point to the application
func main() {
	HandleRequests( )
}
