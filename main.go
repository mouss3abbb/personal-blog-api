package main

import (
	"fmt"
	"net/http"
	"personal-blog-api/controllers"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Welcome to the Personal Blog API")
		fmt.Println("GET Root Endpoint")
	})
	router.HandleFunc("/articles/get/all", controllers.GetAllArticles)
	router.HandleFunc("/articles/get/{id}", controllers.GetArticleById)
	router.HandleFunc("/articles/create", controllers.CreateArticle).Methods("POST")
	router.HandleFunc("/articles/delete/{id}", controllers.DeleteById)
	router.HandleFunc("/articles/update/{id}", controllers.UpdateById)
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		panic(err)
	}
	fmt.Println("Listening at port 8080: ")
}
