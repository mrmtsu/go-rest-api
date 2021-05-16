package main

import (
	"encoding/json"
	"go-rest-api/database"
	"net/http"

	"github.com/gorilla/mux"
)

type Article struct {
	Id int `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
}

func getAllArticles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	articles := []Article{}
	database.DB.Find(&articles)
	json.NewEncoder(w).Encode(articles)
}

func getArticle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	articleID := params["id"]

	article := Article{}
	database.DB.First(&article, articleID)
	json.NewEncoder(w).Encode(article)
}

func createArticle(w http.ResponseWriter, r *http.Request) {
	createArticles := Article{}
	json.NewDecoder(r.Body).Decode(&createArticles)
	database.DB.Create(&createArticles)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(createArticles)
}

func updateArticle(w http.ResponseWriter, r *http.Request) {
	updateArticle := Article{}
	json.NewDecoder(r.Body).Decode(&updateArticle)
	database.DB.Save(&updateArticle)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updateArticle)
}

func deleteArticle(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	articleID := params["id"]
	database.DB.Delete(&Article{}, articleID)
	w.WriteHeader(http.StatusNoContent)
}

func main() {
	database.Connect()
	sqlDB, _ := database.DB.DB()
	defer sqlDB.Close()

	router := mux.NewRouter()

	router.HandleFunc("/articles", getAllArticles).Methods("GET")
	router.HandleFunc("/articles/{id}", getArticle).Methods("GET")
	router.HandleFunc("/articles", createArticle).Methods("POST")
	router.HandleFunc("/articles/{id}", updateArticle).Methods("PUT")
	router.HandleFunc("/articles/{id}", deleteArticle).Methods("DELETE")

	http.ListenAndServe(":8000", router)
}
