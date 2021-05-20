package models

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/dileofrancoj/blog-app/config"
	"github.com/dileofrancoj/blog-app/constants"
	"github.com/dileofrancoj/blog-app/structs"
	"github.com/dileofrancoj/blog-app/utils"
)

var db = config.MongoC.Database(constants.DB_NAME)
var collection = db.Collection(constants.USERS_COLLECTIONS)

func IsValidUser(username string) (structs.User, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15 * time.Second)
	defer cancel()

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

func Auth(username string, password string) (structs.User, bool) {
	user, found := IsValidUser(username)
	if found == false {
		return user, false
	}
	result := utils.CompareHash(password, user.Password)
	
	if result == true {
		return user,true
	}
	return user,false
}

func Register (user structs.User) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15 * time.Second)
	defer cancel()

	pwd, _ := utils.Crypt(user.Password)
	user.Password = pwd
	result, err := collection.InsertOne(ctx,user)
	if err != nil {
		return "", false, err
	}
	ObjID, _ := result.InsertedID.(primitive.ObjectID)
	return ObjID.String(), true, nil

}