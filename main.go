package main

import (
	"log"
	"net/http"

	"github.com/matildez0/simple-go-mod/config"
)

//ponto de entrada da aplicação

func main() {
	dbConnection := config.SetupDataBase()

	defer dbConnection.Close()

	//incializa o servidor e as rotas
	log.Fatal(http.ListenAndServe(":8080", nil))

}
