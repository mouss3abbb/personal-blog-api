package main

import (
	"fmt"
	"net/http"
	"os"
	"personal-blog-api/controllers"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
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
	fmt.Printf("Listening at port %s \n", os.Getenv("API_PORT"))
	err := http.ListenAndServe(":"+os.Getenv("API_PORT"), router)
	if err != nil {
		panic(err)
	}
}
