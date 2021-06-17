package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Article struct {
	Name         string `json:"Name"`
	Designation  string `json:"Designation"`
	Company_name string `json:"Company_name"`
}
type Articles []Article

func allArticles(w http.ResponseWriter, r *http.Request) {
	articles := Articles{
		Article{Name: "Subi", Designation: "Software Trainee", Company_name: "Qss Technosoft"},
	}
	fmt.Println("EndPoint Hit: All Articles Endpoint")
	json.NewEncoder(w).Encode(articles)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Homepage!")
	fmt.Println("Endpoint Hit: homePage")

}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/articles", allArticles)
	log.Fatal(http.ListenAndServe(":4040", nil))

}

func main() {
	handleRequests()
}
