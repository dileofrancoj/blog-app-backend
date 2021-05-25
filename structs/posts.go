package structs
type Post struct {
	ID			string	`bson:"_id,omitempty" json:"id"`
	Title		string	`bson:"title" json:"title"`
	Body    	string  `bson:"body"  json:"body"`
	Visible 	bool 	`bson:"visible,omitempty" json:"visible,omitempty"`
}

type Posts []*Post

/*CreatePost : función encargada de crear un posteo*/
func (p *Post) CreatePost(post Post) (bool, error) {
	return true, nil
}
