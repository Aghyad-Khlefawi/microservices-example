package servicecollection

import "go.mongodb.org/mongo-driver/mongo"

type ServiceCollection struct {
	Configurations map[string]string
	MongoClient *mongo.Client
}

func NewServiceCollection(config map[string]string, mongoClient *mongo.Client) *ServiceCollection {
	return &ServiceCollection{
		config,
		mongoClient,
	}
}