package conceptos

import (
	"context"
	"net/http"
	"github.com/rubengp99/golang-rest-api/server/data"
	"github.com/rubengp99/golang-rest-api/commons"
	"go.mongodb.org/mongo-driver/bson"
	"encoding/json"
)

const model = "adm_conceptos"

// GetAll returns all concepts
func GetAll( res http.ResponseWriter, req *http.Request){
	db, status := data.Connection()
	var cur, err interface{}

	if status.Code == 500{
		json.NewEncoder(res).Encode(commons.NewResponse{0, nil, commons.InternalServerError()})
	}else{
		cur, err := db.Collection(model).Find(context.TODO(), bson.D, findOptions)
	}
	
}