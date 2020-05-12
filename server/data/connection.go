package data


import (
	"fmt"
	"context"
	"time"
	"os"
	//autoload for .env
	_ "github.com/joho/godotenv/autoload"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"github.com/rubengp99/golang-rest-api/commons"
)

// Connection returns a MongoDB Database and a APIResponse for the status of the connection
func Connection() (*mongo.Database, commons.APIResponse){
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(
		"mongodb+srv://"+os.Getenv("MONGO_USER")+":"+os.Getenv("MONGO_PASSWORD")+"@"+os.Getenv("MONGO_CLUSTER")+"-8p1no.mongodb.net/test?retryWrites=true&w=majority",
	))

	err = client.Ping(context.TODO(), nil)

    if err != nil {
		fmt.Println("Error at Mongo-CONNECT err:",err)
		return nil , commons.InternalServerError()
	}
	
	return client.Database(os.Getenv("MONGO_DB")), commons.Ok()
}