package main

import (
	"fmt"
	 _ "github.com/joho/godotenv/autoload"	
	"net/http"
	"os"
	"github.com/gorilla/mux"
	"github.com/rubengp99/golang-rest-api/server/api/conceptos"	
)

func main(){
	router := router()
	fmt.Println("SERVER running on port",os.Getenv("API_PORT"))
	fmt.Println(http.ListenAndServe( ":"+os.Getenv("API_PORT") ,  router))
}

//Router exports the router of the API
func router() *mux.Router{
	Router := mux.NewRouter().StrictSlash(true)
	Router.Handle("/api/conceptos", conceptos.Router())
	
	return Router
}