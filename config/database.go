package config

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*MongoC Me permite la conexión a la base de datos*/
var MongoC = ConnectDatabase()
var DB_URI = "mongodb+srv://francodileo-blog:2Eg4Q3s2rfZsb2g@cluster0.h04ks.mongodb.net/test"

/*ConnectDatabase Me permite la conexión a la base de datos*/
func ConnectDatabase() *mongo.Client {
	clientOptions := options.Client().ApplyURI(DB_URI)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal((err))
		return client
	}

	log.Print("Conectado a la DB")
	return client
}