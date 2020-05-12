package main

import (
	 _ "github.com/joho/godotenv/autoload"	
	"net/http"
	"os"
)

func main(){
	router := router()
	http.ListenAndServe( os.Getenv("API_PORT") , router);
}
