package main

import (
	"go-todo/controllers/todo"
	"go-todo/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()

}

func main() {
	router := gin.Default()
	router.POST("/todo/create", todo.CreateTodo)
	router.GET("/todo/get-all", todo.GetTodos)
	router.GET("/todo/details/:id", todo.GetTodo)
	router.PUT("/todo/update/:id", todo.UpdateTodo)
	router.DELETE("/todo/delete/:id", todo.DeleteTodo)
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "PONG",
		})
	})
	router.Run()
}
