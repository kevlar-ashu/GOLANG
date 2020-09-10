package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Article Structure for multiple data type field..
type Article struct {
	Title   string `json:"Title"`
	Desc    string `json:"Desc"`
	Content string `json:"Content"`
}

// Articles array for storing multiple field...
type Articles []Article

func allArticles(w http.ResponseWriter, r *http.Request) {
	articles := Articles{
		Article{Title: "Test title", Desc: "Test Description", Content: "Hello World"},
		Article{Title: "Test You", Desc: "Test Desc", Content: "My World"},
	}

	fmt.Println("Endpoint Hit: All Articles")
	json.NewEncoder(w).Encode(articles)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage Endpoint Hit")
}

//func testPostArticles(w http.ResponseWriter, r *http.Request) {
//	fmt.Fprintf(w, "test Post Endpoint Worked")
//}

func handleRequests() {

	//myRouter := mux.NewRouter().StrictSlash(true)

	http.HandleFunc("/", homePage)
	http.HandleFunc("/articles", allArticles)
	//myRouter.HandleFunc("/articles", testPostArticles).Methods("POST")
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func main() {
	handleRequests()
}
