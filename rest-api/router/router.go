package router

import (
	"rest-api/controllers"

	"github.com/gorilla/mux"
)

func Init() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/", controllers.HomeGetHandler).Methods("GET")

	router.HandleFunc("/users", controllers.GetUsersHandler).Methods("GET")
	router.HandleFunc("/users/{id}", controllers.GetUserHandler).Methods("GET")
	router.HandleFunc("/users", controllers.PostUserHandler).Methods("POST")
	router.HandleFunc("/users/{id}", controllers.PutUserHandler).Methods("PUT")
	router.HandleFunc("/users/{id}", controllers.DeleteUserHandler).Methods("DELETE")

	return router
}
