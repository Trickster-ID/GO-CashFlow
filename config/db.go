package config

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// func SetUpDatabaseConnection() {
// 	var errcon error
// 	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
// 	_, errcon = mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
// 	if errcon != nil {
// 		panic("error while connect : "+ errcon.Error())
// 	}
// }

var client *mongo.Client

func ResolveClientDB() *mongo.Client {
	if client != nil {
		return client
	}

	var err error
	// TODO add to your .env.yml or .config.yml MONGODB_URI: mongodb://localhost:27017
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err = mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		panic(err)
	}

	// check the connection
	err = client.Ping(context.Background(), nil)
	if err != nil {
		panic(err)
	}

	// TODO optional you can log your connected MongoDB client
	fmt.Println("Database connection success")
	return client
}

func CloseClientDB(db *mongo.Client) {
	if client == nil {
		return
	}

	err := db.Disconnect(context.TODO())
	if err != nil {
		panic(err)
	}
}
