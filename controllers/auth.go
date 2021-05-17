package controllers

import (
	"log"

	"github.com/dileofrancoj/blog-app/models"
	"github.com/dileofrancoj/blog-app/structs"
	"github.com/gin-gonic/gin"
)

func Auth(ctx *gin.Context) {
	var login structs.Login
	err := ctx.ShouldBindJSON(&login)
	if err!=nil {
		log.Fatal(err.Error())
		ctx.JSON(500, gin.H{
			"message" : "Ocurrió un error interno",
		})
		return
	}
	user, _ := models.IsValidUser(login.Username)
	if err!=nil {
		log.Fatal(err.Error())
		ctx.JSON(500, gin.H{
			"message" : "Ocurrió un error interno",
		})
		return
	}

	ctx.JSON(200,gin.H{
		"message" : "El usuario existe y tiene un ID: " + user.ID, 
	})

}