package routes

import (
	"go-rest-api/controllers"
	"net/http"

	"github.com/gorilla/mux"
)

func Setup() {
	router := mux.NewRouter()

	router.HandleFunc("/articles", controllers.GetAllArticles).Methods("GET")
	router.HandleFunc("/articles/{id}", controllers.GetArticle).Methods("GET")
	router.HandleFunc("/articles", controllers.CreateArticle).Methods("POST")
	router.HandleFunc("/articles/{id}", controllers.UpdateArticle).Methods("PUT")
	router.HandleFunc("/articles/{id}", controllers.DeleteArticle).Methods("DELETE")

	http.ListenAndServe(":8000", router)
}
