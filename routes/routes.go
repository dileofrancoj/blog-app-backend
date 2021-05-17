package routes

import (
	"os"

	"github.com/dileofrancoj/blog-app/controllers"
	"github.com/gin-gonic/gin"
)

/*Routes -> handle project routing*/
func Routes(){
	PORT := os.Getenv("PORT")
	Router := gin.Default()
	api := Router.Group("/api")
	{
		api.GET("/posts", controllers.GetPosts)
		api.POST("/posts", controllers.CreatePost)
		api.POST("/auth", controllers.Auth)
		api.POST("/register", controllers.Register)
	}
	Router.Run(":"+PORT)
}