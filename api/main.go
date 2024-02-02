package main

import (
    "net/http"
    "github.com/gorilla/mux"
)

func getUsers(w http.ResponseWriter, r *http.Request) {
    // Lógica para obtener usuarios
}

func createUser(w http.ResponseWriter, r *http.Request) {
    // Lógica para crear un usuario
}

func main() {
    router := mux.NewRouter()

    router.HandleFunc("/users", getUsers).Methods("GET")
    router.HandleFunc("/users", createUser).Methods("POST")

    http.ListenAndServe(":8080", router)
}