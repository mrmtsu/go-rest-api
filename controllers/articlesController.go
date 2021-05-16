package controllers

import (
	"encoding/json"
	"go-rest-api/database"
	"go-rest-api/models"
	"net/http"

	"github.com/gorilla/mux"
)

func GetAllArticles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	articles := []models.Article{}
	database.DB.Find(&articles)
	json.NewEncoder(w).Encode(articles)
}

func GetArticle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	articleID := params["id"]

	article := models.Article{}
	database.DB.First(&article, articleID)
	json.NewEncoder(w).Encode(article)
}

func CreateArticle(w http.ResponseWriter, r *http.Request) {
	createArticles := models.Article{}
	json.NewDecoder(r.Body).Decode(&createArticles)
	database.DB.Create(&createArticles)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(createArticles)
}

func UpdateArticle(w http.ResponseWriter, r *http.Request) {
	updateArticle := models.Article{}
	json.NewDecoder(r.Body).Decode(&updateArticle)
	database.DB.Save(&updateArticle)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updateArticle)
}

func DeleteArticle(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	articleID := params["id"]
	database.DB.Delete(&models.Article{}, articleID)
	w.WriteHeader(http.StatusNoContent)
}
