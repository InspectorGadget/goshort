package main

import (
	"github.com/InspectorGadget/goshort/controllers"
	"github.com/InspectorGadget/goshort/initializers"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()

	if err := initializers.ConnectToDB(); err != nil {
		panic(err)
	}
}

func main() {
	r := gin.Default()

	r.GET("/", controllers.Index)

	r.GET("/users", controllers.ListUsers)
	r.POST("/users", controllers.AddUser)
	r.DELETE("/users/:id", controllers.DeleteUser)
	r.POST("/users/:id/url", controllers.AddUrlToUser)

	if err := r.Run(":3000"); err != nil {
		panic(err)
	}
}
