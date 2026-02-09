package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/matildez0/simple-go-mod/config"
	"github.com/matildez0/simple-go-mod/models"
)

//ponto de entrada da aplicação

func main() {
	dbConnection := config.SetupDataBase()

	_, err := dbConnection.Exec(models.CreateTableSQL)
	if err != nil {
		log.Fatal(err)
	}
	router := mux.NewRouter()
	router.HandleFunc("/tasks", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello, World!"))
	}).Methods("GET")

	defer dbConnection.Close()

	//incializa o servidor e as rotas
	log.Fatal(http.ListenAndServe(":8080", router))

}
