package conceptos

import (
	"fmt"
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
	ctx := context.TODO()
	
	if status.Code == 500{
		json.NewEncoder(res).Encode(commons.NewResponse{0, nil, commons.InternalServerError()})
	}else{
		cursor, err := db.Collection(model).Find(ctx, bson.M{})
		
		var data []bson.M
		if err = cursor.All(ctx, &data); err != nil {
			panic(err)
		}
		//json.NewEncoder(res).Encode(commons.NewResponse{0, nil, commons.InternalServerError()})
		fmt.Println(cursor.All(ctx, &data));
	}
	
}