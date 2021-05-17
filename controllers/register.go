package controllers

import (
	"fmt"
	"log"

	"github.com/dileofrancoj/blog-app/models"
	"github.com/dileofrancoj/blog-app/structs"
	"github.com/gin-gonic/gin"
)

func Register(ctx *gin.Context) {
	var user structs.User
	err := ctx.ShouldBindJSON(&user)
	if err!=nil {
		log.Fatal(err.Error())
		ctx.JSON(500, gin.H{
			"message" : "Ocurrió un error interno",
		})
		return
	}
	_, existUser := models.IsValidUser(user.Username)
	if(existUser) {
		ctx.JSON(200, gin.H{
			"message" : "El usuario ya existe",
		})
		return
	}
	id, registered ,err := models.Register(user)
	fmt.Print(id, registered)
}