package controllers

import (
	"log"

	"github.com/dileofrancoj/blog-app/models"
	"github.com/dileofrancoj/blog-app/structs"
	"github.com/dileofrancoj/blog-app/utils"
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
	user, found := models.Auth(login.Username,login.Password)
	if err!=nil {
		log.Fatal(err.Error())
		ctx.JSON(500, gin.H{
			"message" : "Ocurrió un error interno",
		})
		return
	}


	if found == false {
		ctx.JSON(401,gin.H{
			"message" : "Usuario o contraseña incorrecta",
		})
		return
	}

	jwt, err := utils.JWT(user)

	if err!=nil {
		log.Fatal(err.Error())
		ctx.JSON(500, gin.H{
			"message" : "Ocurrió un error interno",
		})
		return
	}


	ctx.JSON(200,gin.H{
		"jwt" : jwt,
	})

}