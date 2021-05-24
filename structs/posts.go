package structs
type Post struct {
	Title	string	`bson:"title" json:"title"`
	Body    string  `bson:"body"  json:"body"`
}

type Posts []*Post

/*CreatePost : función encargada de crear un posteo*/
func (p *Post) CreatePost(post Post) (bool, error) {
	return true, nil
}
