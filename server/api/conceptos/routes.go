package conceptos

import (
	"github.com/gorilla/mux"
)

// Router returns a Router Object with EndPoint's routes
func Router() *mux.Router{
	conceptos := mux.NewRouter().StrictSlash(true)
	
	conceptos.HandleFunc("/",GetAll).Methods("GET")
	return conceptos
}