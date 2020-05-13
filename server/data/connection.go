package data


import (
	"fmt"
	"context"
	"os"
	//autoload for .env
	_ "github.com/joho/godotenv/autoload"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"github.com/rubengp99/golang-rest-api/commons"
)

// Connection returns a MongoDB Database and a APIResponse for the status of the connection
func Connection() (*mongo.Database, commons.APIResponse){
	ctx := context.Background()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(
		"mongodb+srv://"+os.Getenv("MONGO_USER")+":"+os.Getenv("MONGO_PASSWORD")+"@"+os.Getenv("MONGO_CLUSTER")+"-8p1no.gcp.mongodb.net/test?retryWrites=true&w=majority",
	))

    if err != nil {
		fmt.Println("Error at Mongo-CONNECT err:",err)
		return nil , commons.InternalServerError()
	}else{
		return client.Database(os.Getenv("MONGO_DB")), commons.Ok()
	}
}