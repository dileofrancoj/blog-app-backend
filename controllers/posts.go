package controllers

import (
	"fmt"
	"log"

	"github.com/dileofrancoj/blog-app/models"
	"github.com/gin-gonic/gin"
)

func GetPosts(ctx *gin.Context){
	ctx.JSON(200,gin.H {
		"posts" : "all posts",
	})
}

func CreatePost(ctx *gin.Context) {
	var post models.Post
	err := ctx.ShouldBindJSON(&post)
	fmt.Println(&post)
	if err != nil {
		log.Fatal(err.Error())
		ctx.JSON(400, gin.H{
			"message" : "Ocurrió un error",
		})
		return
	}

}