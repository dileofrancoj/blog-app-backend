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
	router := gin.Default()

	api := router.Group("/api")
	{
		api.GET("/posts", controllers.GetPosts)
		api.GET("/posts/:id", controllers.GetPost)
		api.POST("/auth", controllers.Auth)
		api.POST("/register", controllers.Register)	

		api.Use(middlewares.ValidateJWT())
		{
			api.POST("/posts" ,controllers.CreatePost)
			api.DELETE("/posts/:id", controllers.DeletePost)
			api.PUT("/posts/:id", controllers.UpdatePost)
		}
		

	}
	router.Run(":"+PORT)
}