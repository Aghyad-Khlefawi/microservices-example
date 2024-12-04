package servicecollection

import "go.mongodb.org/mongo-driver/mongo"

type ServiceCollection struct {
	MongoClient *mongo.Client
}

func NewServiceCollection(mongoClient *mongo.Client) *ServiceCollection {
	return &ServiceCollection{
		mongoClient,
	}
}

func SetDefaultServiceCollection(sc *ServiceCollection){
	defaultsc = sc
}

func Default() *ServiceCollection{
	return defaultsc
}

var defaultsc *ServiceCollection;
