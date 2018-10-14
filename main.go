
package main

import (
	"log"
	"net/http"
	"github.com/gorilla/handlers"
	"musicstore/modules/album"
	"github.com/joho/godotenv"
	"os"
	
)

func main(){
	err := godotenv.Load()
  if err != nil {
    log.Fatal("Error loading .env file")
	}
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}
	
	//router call 
	router := album.NewRouter()
	allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "DELETE", "PUT"})
	log.Fatal(http.ListenAndServe(":"+ port, handlers.CORS(allowedOrigins, allowedMethods)(router)))
}
