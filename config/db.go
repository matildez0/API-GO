package config

import (
	"database/sql"
	"log"
	"os"
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

//vai buscar as varáveis do env, inicializa a conexão com a base de dados

func SetupDataBase() *sql.DB{
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_DATABASE")

	connectionStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPassword, dbName)
	dbConnection, err := sql.Open("postgres", connectionStr)

	if err != nil {
		log.Fatal( err)
	}

	err = dbConnection.Ping()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Sucesso na conexão com a base de dados")

	return dbConnection


}