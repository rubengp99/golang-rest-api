package conceptos

import (
	"github.com/gorilla/mux"
)

func conceptosRouter() mux.Router{
	conceptos := mux.NewRouter().StrictSlash(true)
	
	conceptos.HandleFunc("/",GetAll)
	return *conceptos
}