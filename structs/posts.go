package structs

import "time"
type Post struct {
	ID			string		`bson:"_id,omitempty" json:"id"`
	Title		string		`bson:"title" json:"title"`
	Body    	string  	`bson:"body"  json:"body"`
	Visible 	bool 		`bson:"visible,omitempty" json:"visible,omitempty"`
	Date		time.Time 	`bson:"created_at" json:"created_at"`
}

type Posts []*Post

type PostActions interface {
	GetPosts() (*Post, error)
	GetPost(id string) (*Post, error)
	DeletePost(id string) error
}