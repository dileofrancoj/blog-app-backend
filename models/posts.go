package models

import (
	"fmt"

	"github.com/dileofrancoj/blog-app/structs"
)

/*CreatePost --> función para crear posteo*/
func CreatePost(p structs.Post) (bool,error){
	fmt.Print("Entra a crear posteo")
	return true,nil
}