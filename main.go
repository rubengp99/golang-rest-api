package main

import (
	//"fmt"
	"github.com/gorilla/mux"
	"./api/conceptos"
)

func main(){
	conceptos := mux.NewRouter().StrictSlash(true)
	conceptos.Handle("",nil)
}