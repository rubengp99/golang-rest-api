package data

import(
	"net/http"
	"encoding/json"
)

func documentation( res http.ResponseWriter, req *http.Request){
	json.NewEncoder(res).Encode(req)
}