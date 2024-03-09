package main

import (
	"go-api/controllers"
	"go-api/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()
	r.POST("/post", controllers.CreatePost)
	r.GET("/post", controllers.GetPost)
	r.GET("/post/:id", controllers.GetPostById)
	r.PUT("/post/:id", controllers.UpdatePost)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
