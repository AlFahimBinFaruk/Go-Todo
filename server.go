package main

import (
	"go-todo/controllers/todo"
	"go-todo/controllers/user"
	"go-todo/initializers"
	"go-todo/middleware"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()

}

func main() {
	router := gin.Default()
	todoGroup := router.Group("/todo", middleware.RequireAuth)
	{
		todoGroup.POST("/create", todo.CreateTodo)
		todoGroup.GET("/get-all", todo.GetTodos)
		todoGroup.GET("/details/:id", todo.GetTodo)
		todoGroup.PUT("/update/:id", todo.UpdateTodo)
		todoGroup.DELETE("/delete/:id", todo.DeleteTodo)
	}
	userGroup := router.Group("/user")
	{
		userGroup.POST("/login", user.LoginUser)
		userGroup.POST("/register", user.RegisterUser)
	}
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "PONG",
		})
	})
	router.Run()
}
