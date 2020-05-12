package main

import (
	"github.com/gorilla/mux"
	"github.com/rubengp99/golang-rest-api/server/api/conceptos"	
)

func router() *mux.Router{
	router := mux.NewRouter().StrictSlash(true)
	router.Handle("api/conceptos", conceptos.Router()).Methods("GET")
	
	return router
}