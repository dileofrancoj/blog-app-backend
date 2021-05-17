package structs

import "fmt"

type Post struct {
	Title	string	`bson:"title" json:"title"`
	Body    string  `bson:"body"  json:"body"`
}

/*CreatePost : función encargada de crear un posteo*/
func (p *Post) CreateUser(post Post) (bool, error) {
	fmt.Println(p)
	return true, nil
}
