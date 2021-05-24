package models

import (
	"context"
	"time"

	"github.com/dileofrancoj/blog-app/config"
	"github.com/dileofrancoj/blog-app/constants"
	"github.com/dileofrancoj/blog-app/structs"
	"go.mongodb.org/mongo-driver/bson"
)

/*CreatePost --> función para crear posteo*/
func CreatePost(p structs.Post) (bool,error){
	return true,nil
}

func GetPosts() (structs.Posts, error) {
	var db = config.MongoC.Database(constants.DB_NAME)
	var collection = db.Collection(constants.POSTS_COLLECTION)

	ctx, cancel := context.WithTimeout(context.Background(), 15 * time.Second)
	defer cancel()
	var posts structs.Posts

	filter := bson.D{}
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