package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"personal-blog-api/models"

	"github.com/gorilla/mux"
)

func GetAllArticles(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(models.Articles)
}

func GetArticleById(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	if id == "" {
		fmt.Fprintln(w, "Error no such id found")
		return
	}
	for _, a := range models.Articles {
		if a.ID == id {
			json.NewEncoder(w).Encode(a)
			return
		}
	}
	fmt.Fprintln(w, "Error no such id found")
}

func CreateArticle(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Error in the body: %v", err.Error())
		return
	}
	var article models.Article
	err = json.Unmarshal(body, &article)
	if err != nil {
		fmt.Fprintf(w, "Error in the body: %v", err.Error())
		return
	}
	models.Articles = append(models.Articles, article)
	json.NewEncoder(w).Encode(article)
}

func DeleteById(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	if id == "" {
		if id == "" {
			fmt.Fprintln(w, "Error no such id found")
			return
		}
	}
	for key, value := range models.Articles {
		if value.ID == id {
			models.Articles = append(models.Articles[:key], models.Articles[key+1:]...)
			json.NewEncoder(w).Encode(value)
			return
		}
	}
	fmt.Fprintln(w, "Error no such id found")
}

func UpdateById(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintln(w, "No appropriate body provided")
		return
	}
	var res map[string]string
	json.Unmarshal(body, &res)
	for i, a := range models.Articles {
		if a.ID == mux.Vars(r)["id"] {
			if title, ok := res["title"]; ok {
				models.Articles[i].Title = title
			}
			if content, ok := res["content"]; ok {
				models.Articles[i].Content = content
			}
			fmt.Fprintln(w, a)
			return
		}
	}
	fmt.Fprintln(w, "Error in the provided body")
}
