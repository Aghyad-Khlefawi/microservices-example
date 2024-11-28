package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/aghyad-khlefawi/identity/api"
	"github.com/aghyad-khlefawi/identity/pkg/servicecollection"
	"github.com/aghyad-khlefawi/identity/utils"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
)

func main() {

	log.Println("Service starting up")


	config := loadConfig()
	sc:= configureServices(config)
	

	startGrpcServer()
	startRestApi(sc)

}

func startGrpcServer(){
	gs:=grpc.NewServer()
	lis,err:= net.Listen("tcp",":5001")
	if err!=nil{
		utils.LogFatalError("Failed to start GRPC server",err)
	}
	go func(){
		fmt.Println("GRPC Server running on port 5001")
		err= gs.Serve(lis)
		if err!=nil{
			utils.LogFatalError("Failed to start GRPC server",err)
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
