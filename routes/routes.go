package routes

import (
	"os"

	"github.com/dileofrancoj/blog-app/controllers"
	"github.com/dileofrancoj/blog-app/middlewares"
	"github.com/gin-gonic/gin"
)

/*Routes -> handle project routing*/
func Routes(){
	PORT := os.Getenv("PORT")
	Router := gin.Default()
	api := Router.Group("/api")
	{

		api.POST("/posts", middlewares.ValidateJWT() ,controllers.CreatePost)
		api.GET("/posts", controllers.GetPosts)
		api.POST("/auth", controllers.Auth)
		api.POST("/register", controllers.Register)
	}
	Router.Run(":"+PORT)
}