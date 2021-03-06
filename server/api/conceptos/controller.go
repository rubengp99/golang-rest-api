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
	ctx := context.Background()
	if status.Code == 500{
		json.NewEncoder(res).Encode(commons.NewResponse{0, nil, commons.InternalServerError()})
	}else{
		cursor, err := db.Collection(model).Find(ctx, bson.M{})

		if err != nil {
			fmt.Println(`Could not find anything at collection `+model+`, error:`,err)
			json.NewEncoder(res).Encode(commons.NewResponse{0, nil, commons.InternalServerError()})
			return
		}

		var data []bson.M
		if err = cursor.All(ctx, &data); err != nil {
			fmt.Println(`Could not access to data at controller `+model+`, error:`,err)
			json.NewEncoder(res).Encode(commons.NewResponse{0, nil, commons.InternalServerError()})
			return
		}

		count := len(data)
		var response commons.APIResponse

		if count == 0 {
			response = commons.Empty()
		}else{
			response = commons.Ok()
		}

		json.NewEncoder(res).Encode(commons.NewResponse{count, data, response})
	}
	
}