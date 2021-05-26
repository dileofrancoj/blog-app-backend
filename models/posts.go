package models

import (
	"context"
	"time"

	"github.com/dileofrancoj/blog-app/config"
	"github.com/dileofrancoj/blog-app/constants"
	"github.com/dileofrancoj/blog-app/structs"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*CreatePost --> función para crear posteo*/
func CreatePost(p structs.Post) (string,error){
	
	db := config.MongoC.Database(constants.DB_NAME)
	collection := db.Collection(constants.POSTS_COLLECTION)

	ctx, cancel := context.WithTimeout(context.Background() , 15 * time.Second)
	defer cancel()

	result, err := collection.InsertOne(
		ctx,
		bson.M{
			"title":p.Title,
			"body": p.Body,
			"visible": true,
		},
	)

	if err != nil {
		return "", err
	}
	ObjId, _ := result.InsertedID.(primitive.ObjectID)
	return ObjId.String(), nil


}

/*Implementacion interna*/
func getPost(id string) {}

func GetPost(id string) (structs.Post,error) {
	var post structs.Post

	var db = config.MongoC.Database(constants.DB_NAME)
	var collection = db.Collection(constants.POSTS_COLLECTION)

	ctx, cancel := context.WithTimeout(context.Background(), 15 * time.Second)
	defer cancel()

	objID,_ := primitive.ObjectIDFromHex(id)
	condition := bson.M{
		"_id" : objID,
		"visible":true,
	}
	err := collection.FindOne(ctx,condition).Decode(&post)
	if err != nil {
		return post, err
	}
	return post, nil

}

func UpdatePost(id string , post structs.Post) error {
	ctx, cancel := context.WithTimeout(context.Background(), 15 * time.Second)
	defer cancel()

	db := config.MongoC.Database(constants.DB_NAME)
	collection := db.Collection(constants.POSTS_COLLECTION)

	objID, _ := primitive.ObjectIDFromHex(id)
	update := bson.M{
		"$set" : bson.M{
			"title":post.Title,
			"body":post.Body,
		},
	}
	condition := bson.M{
		"_id" : objID,
	}
	_, err:= collection.UpdateOne(ctx,condition,update)
	return err
}

func DeletePost(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 15 * time.Second)
	defer cancel()

	var db = config.MongoC.Database(constants.DB_NAME)
	collection := db.Collection(constants.POSTS_COLLECTION)

	objID, _ := primitive.ObjectIDFromHex(id)
	condition := bson.M{
		"_id" : objID,
	}

	_,err := collection.DeleteOne(ctx,condition)

	return err

}

func GetPosts() (structs.Posts, error) {
	var db = config.MongoC.Database(constants.DB_NAME)
	var collection = db.Collection(constants.POSTS_COLLECTION)

	ctx, cancel := context.WithTimeout(context.Background(), 15 * time.Second)
	defer cancel()
	var posts structs.Posts

	filter := bson.M{"visible":true}
	cur,err := collection.Find(ctx,filter)

	if err != nil {
		return posts,err
	}

	for cur.Next(ctx) {
		var post structs.Post
		err := cur.Decode(&post)
		if err != nil {
			return posts, err
		}

		posts = append(posts, &post)
	}

	return posts,nil
}