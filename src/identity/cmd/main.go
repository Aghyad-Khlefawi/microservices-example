package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/aghyad-khlefawi/identity/api"
	"github.com/aghyad-khlefawi/identity/pkg/servicecollection"
	"github.com/aghyad-khlefawi/identity/utils"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"

	pb "github.com/aghyad-khlefawi/identity/pkg/grpc"
)

func main() {

	log.Println("Service starting up")

	loadConfig()
	sc := configureServices()

	startGrpcServer()
	startRestApi(sc)

}

func startGrpcServer() {
	gs := grpc.NewServer()
	pb.RegisterIdentityServiceServer(gs, pb.NewIdentityService())

	lis, err := net.Listen("tcp", ":5001")
	if err != nil {
		utils.LogFatalError("Failed to start GRPC server", err)
	}
	go func() {
		fmt.Println("GRPC Server running on port 5001")
		err = gs.Serve(lis)
		if err != nil {
			utils.LogFatalError("Failed to start GRPC server", err)
		}
	}()
}

func configureServices() *servicecollection.ServiceCollection {

	log.Println("Configuring application services")

	//DB configuration
	dbConnection,ok := os.LookupEnv("DbConnection")
	if !ok {
		utils.LogFatal("Couldn't find the database connection string in the configurations")
	}

	log.Println("Conneting to database")
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(dbConnection))

	if err != nil && os.Getenv("IgnoreDbError") != "true"{
		utils.LogFatalError("Couldn't connect to the database", err)
	}

	sc := servicecollection.NewServiceCollection(client)
	servicecollection.SetDefaultServiceCollection(sc)
	return sc
}

func loadConfig() {

	log.Println("Reading configurations")

	err := godotenv.Load()

	if err != nil {
		log.Print("Failed to load .env file")
	} 
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
