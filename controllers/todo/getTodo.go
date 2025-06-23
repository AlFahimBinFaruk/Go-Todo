package todo

import (
	"go-todo/initializers"
	"go-todo/models"

	"github.com/gin-gonic/gin"
)

func GetTodo(c *gin.Context) {
	id := c.Param("id")
	var todo models.TODO
	result := initializers.DB.First(&todo, id)
	if result.Error != nil {
		c.JSON(400, gin.H{
			"msg": "failed to get todo!!",
		})
		return
	}
	c.JSON(200, gin.H{
		"todo": todo,
	})
}
