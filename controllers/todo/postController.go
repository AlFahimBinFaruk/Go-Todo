package todo

import (
	"go-todo/initializers"
	"go-todo/models"

	"github.com/gin-gonic/gin"
)

func CreateTodo(c *gin.Context) {
	type REQ struct {
		Title string
		Desc  string
	}
	var body REQ
	c.Bind(&body)
	todo := models.TODO{Title: body.Title, Desc: body.Desc}
	result := initializers.DB.Create(&todo)
	if result.Error != nil {
		c.JSON(400, gin.H{
			"msg": "Failed to add Todo.",
		})
		return
	}
	c.JSON(202, gin.H{
		"todo": todo,
	})
}
