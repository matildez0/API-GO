package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/matildez0/simple-go-mod/config"
	"github.com/matildez0/simple-go-mod/handlers"
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

	taskHandler := handlers.NewTaskHandler(dbConnection)

	router.HandleFunc("/tasks", taskHandler.ReadTasks).Methods("GET")
	router.HandleFunc("/tasks", taskHandler.CreateTask).Methods("POST")
	router.HandleFunc("/tasks/{id}", taskHandler.UpdateTask).Methods("PUT")
	router.HandleFunc("/tasks/{id}", taskHandler.DeleteTask).Methods("DELETE")
	/*router.HandleFunc("/tasks", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello, World!"))
	}.Methods("GET")*/

	defer dbConnection.Close()

	//incializa o servidor e as rotas
	log.Fatal(http.ListenAndServe(":8080", router))

}
