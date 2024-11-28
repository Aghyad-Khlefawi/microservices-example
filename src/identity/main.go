package main

import (
	"context"
	"fmt"
	"log"

	"github.com/aghyad-khlefawi/identity/api"
	"github.com/aghyad-khlefawi/identity/pkg/servicecollection"
	"github.com/aghyad-khlefawi/identity/utils"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {

	log.Println("Service starting up")


	config := loadConfig()
	sc:= configureServices(config)
	
	startRestApi(sc)


	defer func(){
		if err:= sc.MongoClient.Disconnect(context.TODO()); err!=nil{
			utils.LogFatalError("Couldn't close the database connection", err)
		}
	}()
}

func configureServices(config map[string] string) *servicecollection.ServiceCollection{
  
  log.Println("Configuring application services")

	//DB configuration
	dbConnection,ok:= config["DbConnection"]
	if !ok {
		utils.LogFatal("Couldn't find the database connection string in the configurations")
	}

	log.Println("Conneting to database")
	client, err:= mongo.Connect(context.TODO(),options.Client().ApplyURI(dbConnection))

	if err!=nil{
		utils.LogFatalError("Couldn't connecto to the database", err)
	}

	return servicecollection.NewServiceCollection(config,client)
}


func loadConfig() map[string]string {

	log.Println("Reading configurations")

	err := godotenv.Load()

	if err != nil {
		utils.LogFatalError("Failed to load env file", err)
	}

	config, err := godotenv.Read()
	if err != nil {
		utils.LogFatalError("Failed to read env file", err)
	}

	return config
}

func startRestApi(sc *servicecollection.ServiceCollection) {
	
	router := gin.Default()

	fmt.Println("HTTP Server listening on port 8080")

	api.RegisterRoutes(router, sc)

	err := router.Run(":8080")

	if err != nil {
		panic(err.Error())
	}

}
