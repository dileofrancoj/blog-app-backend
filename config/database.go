package config

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*MongoC Me permite la conexión a la base de datos*/
var MongoC = ConnectDatabase()
var clientOptions = options.Client().ApplyURI("mongodb+srv://francodileo-blog:2Eg4Q3s2rfZsb2g@cluster0.h04ks.mongodb.net/test")

/*ConnectDatabase Me permite la conexión a la base de datos*/
func ConnectDatabase() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal((err))
		return client
	}

	log.Printf("Conectado a la base de datos")
	return client
}