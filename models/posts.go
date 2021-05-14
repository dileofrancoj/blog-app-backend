package models

import "fmt"

type Post struct {
	Title	string	`bson:"title" json:"title,omitempty"`
	Body    string  `bson:"body"  json:"body,omitempty"`
}

/*CreatePost : función encargada de crear un posteo*/
func (p *Post) CreateUser(post Post) (bool, error) {
	fmt.Println(p)
	return true, nil
}

type Actions interface {
	CreatePost(pst Post) (bool, error)
}