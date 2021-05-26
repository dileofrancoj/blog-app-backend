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
	ObjId,error := models.CreatePost(post)
	if error !=nil {
		log.Fatal(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message" : "Ocurrió un error al crear el posteo",
		})
		return	
	}

	ctx.JSON(http.StatusCreated, gin.H{
			"message" : "Posteo creado correctamente",
			"id": ObjId,
		})
}


func DeletePost(c *gin.Context) {
	id := c.Param("id")
	err := models.DeletePost(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H {
			"message" : "Ocurrió un error interno",
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"message" : "Posteo eliminado",
		"id" : id,
	})

}

func UpdatePost(c *gin.Context){
	id := c.Param("id")
	var post structs.Post
	err := c.ShouldBindJSON(&post)

	if err != nil {
		log.Fatal(err.Error())
		c.JSON(500, gin.H{
			"message" : "El formato de datos es inválido",
		})
		return
	}

	postErr := models.UpdatePost(id,post)
	if postErr != nil {
		c.JSON(http.StatusInternalServerError,gin.H{
			"message" : "Ocurrió un error al actualizar el post",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message":"Posteo actualizado",
		"post":post,
	})

}