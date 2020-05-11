package data

import (
	"os"
	"database/sql"
	//driver for mysql
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)


func connect(url string){
	godotenv.Load(".env")
	connection, _ := sql.Open("mysql", os.Getenv("username")+":"+os.Getenv("password")+"@/"+os.Getenv("database"))
	DoQuery(url)
	connection.Close()
}