package todo

import (
	"go-todo/initializers"
	"go-todo/models"

	"github.com/gin-gonic/gin"
)

func GetTodos(c *gin.Context) {
	var todos []models.TODO
	result := initializers.DB.Find(&todos)
	if result.Error != nil {
		c.JSON(400, gin.H{
			"msg": "failed to get todo!!",
		})
		return
	}
	c.JSON(200, gin.H{
		"todos": todos,
	})
}
