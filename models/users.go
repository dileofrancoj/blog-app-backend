package models

import (
	"context"
	"fmt"
	"time"

	"github.com/dileofrancoj/blog-app/config"
	"github.com/dileofrancoj/blog-app/structs"
	"go.mongodb.org/mongo-driver/bson"
)

func IsValidUser(username string) (structs.User, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15 * time.Second)
	defer cancel()

	db := config.MongoC.Database("blog")
	collection := db.Collection("users")
	condition := bson.M{"username":username}
	var result structs.User
	err := collection.FindOne(ctx,condition).Decode(&result)
	ID := result.ID
	fmt.Print(ID)
	if err != nil {
		return result, false
	}
	return result, true

}