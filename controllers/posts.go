package controllers

import (
	"fmt"
	"log"

	"github.com/dileofrancoj/blog-app/models"
	"github.com/dileofrancoj/blog-app/structs"
	"github.com/gin-gonic/gin"
)

func GetPosts(ctx *gin.Context){
	ctx.JSON(200,gin.H {
		"posts" : "all posts",
	})
}

func CreatePost(ctx *gin.Context) {
	var post structs.Post
	err := ctx.ShouldBindJSON(&post)
	fmt.Println(&post)
	if err!=nil {
		log.Fatal(err.Error())
		ctx.JSON(400, gin.H{
			"message" : "Ocurrió un error",
		})
		return
	}
	created,error := models.CreatePost(post)
	if error !=nil {
		log.Fatal(err.Error())
		ctx.JSON(400, gin.H{
			"message" : "Ocurrió un error al crear el posteo",
		})
		return	
	}
	if created == true {
		ctx.JSON(200, gin.H{
			"message" : "Posteo creado correctamente",
		})
		return	
	}

}