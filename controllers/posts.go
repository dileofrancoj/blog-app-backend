package controllers

import (
	"log"
	"net/http"

	"github.com/dileofrancoj/blog-app/models"
	"github.com/dileofrancoj/blog-app/structs"
	"github.com/gin-gonic/gin"
)

func GetPosts(ctx *gin.Context){

	posts, err := models.GetPosts()
	if err != nil {
		ctx.JSON(500,gin.H {
			"message" : "Ocurrió un error",
		})

	}
	ctx.JSON(http.StatusOK,gin.H{
		"posts":posts,
	})
}

func GetPost(ctx *gin.Context) {
	id := ctx.Param("id")
	post, err := models.GetPost(id)
	if err!=nil {
		log.Fatal(err.Error())
		ctx.JSON(500, gin.H{
			"message" : "Ocurrió un error interno",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"post" : post,
	})
}

func CreatePost(ctx *gin.Context) {
	var post structs.Post
	err := ctx.ShouldBindJSON(&post)
	if err!=nil {
		log.Fatal(err.Error())
		ctx.JSON(500, gin.H{
			"message" : "Ocurrió un error interno",
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